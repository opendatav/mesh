/*
 * Copyright (c) 2019, 2023, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package proxy

import (
	"context"
	"fmt"
	"github.com/opendatav/mesh/client/golang/log"
	"github.com/opendatav/mesh/client/golang/mpc"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	tcpmiddleware "github.com/traefik/traefik/v3/pkg/server/middleware/tcp"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"time"
)

func init() {
	tcpmiddleware.Provide(new(tcpAccessLog))
	tcp.AsContextProvider(new(contextProvider))
}

type contextProvider struct {
}

func (that *contextProvider) New() context.Context {
	return mpc.Context()
}

func (that *contextProvider) Set(ctx context.Context, key string, value any) {
	if nil == value {
		return
	}
	if x, ok := ctx.(prsim.Context); ok {
		x.GetAttachments()[key] = fmt.Sprintf("%v", value)
	}
}

func (that *contextProvider) Get(ctx context.Context, key string) any {
	if x, ok := ctx.(prsim.Context); ok {
		return x.GetAttachments()[key]
	}
	return ""
}

type tcpAccessLog struct {
}

func (that *tcpAccessLog) Name() string {
	return "tcp-accesslog"
}

func (that *tcpAccessLog) Priority() int {
	return 0
}

func (that *tcpAccessLog) Scope() int {
	return 0
}

func (that *tcpAccessLog) New(ctx context.Context, next tcp.Handler, name string) (tcp.Handler, error) {
	return &tcpAccessLogHandler{name: name, next: next}, nil
}

type tcpAccessLogHandler struct {
	name string
	next tcp.Handler
}

func (that *tcpAccessLogHandler) ServeTCP(conn tcp.WriteCloser) {
	timestamp := time.Now().UnixMilli()
	that.next.ServeTCP(conn)
	if x, ok := conn.Context().(prsim.Context); ok {
		log.Info(conn.Context(), "digest,TCP,%s,%s,%s,%s,%s,%s,%dms,%d,%d,%s",
			conn.RemoteAddr(),
			conn.LocalAddr(),
			x.GetAttachment("RequestClientAddr"),
			x.GetAttachment(accesslog.ServiceAddr),
			x.GetAttachment(accesslog.ServiceName),
			x.GetAttachment(accesslog.RequestHost),
			time.Now().UnixMilli()-timestamp,
			0,
			0,
			"Y")
	}
}

//
//type tcpSNIRateLimiters struct {
//}
//
//func (that *tcpSNIRateLimiters) Name() string {
//	return tcp.DefaultName
//}
//
//func (that *tcpSNIRateLimiters) New(proxy string, serverName string, dialer px.Dialer) px.Dialer {
//	if "" == serverName || !types.MatchURNDomain(serverName) {
//		return &tcpSNIRateLimiter{serverName: serverName, nodeId: "", dialer: dialer}
//	}
//	names := strings.Split(serverName, ".")
//	if len(names) < 3 {
//		return &tcpSNIRateLimiter{serverName: serverName, nodeId: "", dialer: dialer}
//	}
//	return &tcpSNIRateLimiter{serverName: serverName, nodeId: names[len(names)-3], dialer: dialer}
//}
//
//type tcpSNIRateLimiter struct {
//	nodeId     string
//	serverName string
//	dialer     px.Dialer
//}
//
//func (that *tcpSNIRateLimiter) Dial(network, addr string) (net.Conn, error) {
//	if !proxy.RateLimitEnable {
//		return that.dialer.Dial(network, addr)
//	}
//
//	stream := limiters.MatchStreamLimiter(strings.ToUpper(that.nodeId))
//	if nil == stream {
//		log.Debug0("TCP connect out %s without rate limit", that.nodeId)
//		return that.dialer.Dial(network, addr)
//	}
//	log.Debug0("TCP connect out %s(%dMbps/s-%dMbps/s)", that.nodeId, stream.upstream, stream.downstream)
//	conn, err := that.dialer.Dial(network, addr)
//	if nil != err {
//		return nil, err
//	}
//	return &rateLimitConnection{
//		conn: conn,
//		lr:   stream.Reader(conn),
//		lw:   stream.Writer(conn),
//		closeWrite: func() error {
//			if c, ok := conn.(tcp.WriteCloser); ok {
//				return c.CloseWrite()
//			}
//			return nil
//		},
//	}, nil
//}
