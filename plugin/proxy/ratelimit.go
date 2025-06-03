/*
 * Copyright (c) 2000, 2099, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package proxy

import (
	"context"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/dsa"
	"github.com/opendatav/mesh/client/golang/log"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/mpc"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/types"
	"github.com/opendatav/mesh/plugin/proxy/tbucket"
	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	"github.com/traefik/traefik/v3/pkg/server/dialer"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
	tx "github.com/traefik/traefik/v3/pkg/server/router/tcp"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"io"
	"math"
	"net"
	"net/http"
	"strings"
	"time"
)

func init() {
	var _ tcp.WriteCloser = new(rateLimitConnection)
	var _ net.Conn = new(rateLimitConnection)
	var _ prsim.Listener = limiters
	tx.Provide(new(tcpRateLimiter))
	dialer.Provide(new(httpRateLimiters))
	macro.Provide(prsim.IListener, limiters)
	middleware.Provide(new(httpRateLimiters))
}

const safeBandwidth = 1024 * 1024 * 1 / 8

var limiters = &rateLimiters{streams: dsa.NewStringMap[*limitStream]()}

type httpRateLimiter struct {
	serverName string
	nodeId     string
	dialer     dialer.Dialer
}

func (that *httpRateLimiter) Dial(network, addr string) (c net.Conn, err error) {
	return that.DialContext(context.Background(), network, addr)
}

func (that *httpRateLimiter) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	if !proxy.RateLimitEnable {
		return that.dialer.DialContext(ctx, network, address)
	}

	stream := limiters.MatchStreamLimiter(strings.ToUpper(that.nodeId))
	if nil == stream {
		log.Debug(ctx, "HTTP connect out %s without rate limit", that.nodeId)
		return that.dialer.DialContext(ctx, network, address)
	}
	log.Debug(ctx, "HTTP connect out %s(%dMbps/s-%dMbps/s)", that.nodeId, stream.upstream, stream.downstream)
	conn, err := that.dialer.DialContext(ctx, network, address)
	if nil != err {
		return nil, err
	}
	return &rateLimitConnection{
		conn: conn,
		lr:   stream.Reader(conn),
		lw:   stream.Writer(conn),
		closeWrite: func() error {
			if c, ok := conn.(tcp.WriteCloser); ok {
				return c.CloseWrite()
			}
			return nil
		},
	}, nil
}

type httpRateLimiterX struct {
	next http.Handler
	name string
}

func (that *httpRateLimiterX) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !proxy.RateLimitEnable {
		that.next.ServeHTTP(w, r)
		return
	}
	sn := prsim.MeshFromNodeId.GetHeader(r.Header)
	stream := limiters.MatchStreamLimiter(strings.ToUpper(sn))
	if nil == stream {
		that.next.ServeHTTP(w, r)
		return
	}
	if _, ok := stream.requestsLimiter.TakeMaxDuration(1, time.Second*5); ok {
		that.next.ServeHTTP(w, r)
		return
	}
	w.WriteHeader(http.StatusTooManyRequests)
	if _, err := w.Write([]byte(cause.Ratelimit.Code)); nil != err {
		log.Warn(mpc.HTTPTracerContext(r), "%s", err)
	}
}

type httpRateLimiters struct {
}

func (that *httpRateLimiters) Name() string {
	return "httpRateLimiters"
}

func (that *httpRateLimiters) Scope() int {
	return 0
}

func (that *httpRateLimiters) New(ctx context.Context, next http.Handler, name string) (http.Handler, error) {
	return &httpRateLimiterX{next: next, name: name}, nil
}

func (that *httpRateLimiters) Priority() int {
	return math.MinInt + 100
}

func (that *httpRateLimiters) Match(ctx context.Context, option dialer.Option) bool {
	return true
}

func (that *httpRateLimiters) Next(ctx context.Context, option dialer.Option, dialer dialer.Dialer) dialer.Dialer {
	if "" == option.ServerName() || !types.MatchURNDomain(option.ServerName()) {
		return &httpRateLimiter{serverName: option.ServerName(), nodeId: "", dialer: dialer}
	}
	names := strings.Split(option.ServerName(), ".")
	if len(names) < 3 {
		return &httpRateLimiter{serverName: option.ServerName(), nodeId: "", dialer: dialer}
	}
	return &httpRateLimiter{serverName: option.ServerName(), nodeId: names[len(names)-3], dialer: dialer}
}

func (that *httpRateLimiters) Proxy(ctx context.Context, option dialer.Option, proxy dialer.Proxy) dialer.Proxy {
	return proxy
}

type tcpRateLimiter struct {
	hello tx.Hello
	conn  tcp.WriteCloser
}

func (that *tcpRateLimiter) Name() string {
	return tx.DefaultName
}

func (that *tcpRateLimiter) Accept(hello tx.Hello, conn tcp.WriteCloser) tcp.WriteCloser {
	tcp.SetContextState(conn.Context(), accesslog.RequestHost, hello.ServerName())
	if !proxy.RateLimitEnable {
		return conn
	}
	if stream := limiters.MatchStreamLimiter(conn.RemoteAddr().String()); nil != stream {
		log.Info0("TCP connect in %s(%dMbps/s)-(%dMbps/s)[%s-%s]", conn.RemoteAddr().String(), stream.upstream, stream.downstream, conn.LocalAddr().String(), conn.RemoteAddr().String())
		return &rateLimitConnection{
			conn:       conn,
			lr:         stream.Reader(conn),
			lw:         stream.Writer(conn),
			closeWrite: conn.CloseWrite,
		}
	}
	if !types.MatchURNDomain(hello.ServerName()) {
		return conn
	}
	names := strings.Split(hello.ServerName(), ".")
	if len(names) < 4 {
		return conn
	}
	nodeId := names[len(names)-4]
	stream := limiters.MatchStreamLimiter(strings.ToUpper(nodeId))
	if nil == stream {
		return conn
	}
	log.Info0("TCP connect in %s(%dMbps/s)-(%dMbps/s)[%s-%s]", nodeId, stream.upstream, stream.downstream, conn.LocalAddr().String(), conn.RemoteAddr().String())
	return &rateLimitConnection{
		conn:       conn,
		lr:         stream.Reader(conn),
		lw:         stream.Writer(conn),
		closeWrite: conn.CloseWrite,
	}
}

type rateLimiters struct {
	streams dsa.Map[string, *limitStream]
}

func (that *rateLimiters) Att() *macro.Att {
	return &macro.Att{Name: "mesh.plugin.proxy.limiter"}
}

func (that *rateLimiters) Btt() []*macro.Btt {
	return []*macro.Btt{prsim.NetworkRouteRefresh}
}

func (that *rateLimiters) Listen(ctx context.Context, event *types.Event) error {
	var routes []*types.Route
	if err := event.TryGetObject(&routes); nil != err {
		return cause.Error(err)
	}
	for _, route := range routes {
		bandwidthUp := 1024 * 1024 * route.Upstream / 8
		bandwidthDown := 1024 * 1024 * route.Downstream / 8
		stream := that.MatchStreamLimiter(strings.ToUpper(route.NodeId))
		if nil != stream && !stream.Diff(route) {
			continue
		}
		log.Info(ctx, "Ratelimit %s upstream:%dMbps#%s downstream:%dMbps#%s. ", route.NodeId, route.Upstream, route.URC().String(), route.Downstream, route.StaticIP)
		if nil == stream {
			xstream := &limitStream{
				nodeId:          strings.ToUpper(route.NodeId),
				upstream:        route.Upstream,
				downstream:      route.Downstream,
				readLimiter:     tbucket.NewFreshRateLimiter(float64(bandwidthDown), bandwidthDown, safeBandwidth, safeBandwidth),
				writeLimiter:    tbucket.NewFreshRateLimiter(float64(bandwidthUp), bandwidthUp, safeBandwidth, safeBandwidth),
				requestsLimiter: tbucket.NewFreshRateLimiter(float64(route.Requests), int64(route.Requests), 99999, 99999),
			}
			that.streams.Put(strings.ToUpper(route.NodeId), xstream)
			that.streams.Put(strings.ToUpper(route.InstId), xstream)
			that.streams.Put(route.StaticIP, xstream)
			that.streams.Put(route.PublicIP, xstream)
			continue
		}
		stream.Refresh(bandwidthUp, bandwidthDown, route)
	}
	return nil
}

func (that *rateLimiters) MatchStreamLimiter(ip string) *limitStream {
	s, ok := that.streams.Get(ip)
	if !ok {
		return nil
	}
	return s
}

type limitStream struct {
	nodeId          string
	upstream        int64
	downstream      int64
	readLimiter     tbucket.FreshRateLimiter
	writeLimiter    tbucket.FreshRateLimiter
	requestsLimiter tbucket.FreshRateLimiter
}

func (that *limitStream) Diff(route *types.Route) bool {
	return that.upstream != route.Upstream || that.downstream != route.Downstream
}

func (that *limitStream) Refresh(bandwidthUp int64, bandwidthDown int64, route *types.Route) {
	that.nodeId = route.NodeId
	that.upstream = route.Upstream
	that.downstream = route.Downstream
	that.readLimiter.Refresh(float64(bandwidthDown), bandwidthDown, safeBandwidth, safeBandwidth)
	that.writeLimiter.Refresh(float64(bandwidthUp), bandwidthUp, safeBandwidth, safeBandwidth)
	that.requestsLimiter.Refresh(float64(route.Requests), int64(route.Requests), 99999, 99999)
}

func (that *limitStream) Reader(reader io.Reader) io.Reader {
	return tbucket.Reader(reader, that.readLimiter)
}

func (that *limitStream) Writer(writer io.Writer) io.Writer {
	return tbucket.Writer(writer, that.writeLimiter)
}

type rateLimitConnection struct {
	conn       net.Conn
	lr         io.Reader
	lw         io.Writer
	closeWrite func() error
}

func (that *rateLimitConnection) Context() context.Context {
	if swc, ok := that.conn.(tcp.WriteCloser); ok {
		return swc.Context()
	}
	return mpc.Context()
}

func (that *rateLimitConnection) CloseWrite() error {
	if nil != that.closeWrite {
		return that.closeWrite()
	}
	return nil
}

func (that *rateLimitConnection) Read(b []byte) (int, error) {
	return that.lr.Read(b)
}

func (that *rateLimitConnection) Write(b []byte) (int, error) {
	return that.lw.Write(b)
}

func (that *rateLimitConnection) Close() error {
	return that.conn.Close()
}

func (that *rateLimitConnection) LocalAddr() net.Addr {
	return that.conn.LocalAddr()
}

func (that *rateLimitConnection) RemoteAddr() net.Addr {
	return that.conn.RemoteAddr()
}

func (that *rateLimitConnection) SetDeadline(t time.Time) error {
	return that.conn.SetDeadline(t)
}

func (that *rateLimitConnection) SetReadDeadline(t time.Time) error {
	return that.conn.SetReadDeadline(t)
}

func (that *rateLimitConnection) SetWriteDeadline(t time.Time) error {
	return that.conn.SetWriteDeadline(t)
}
