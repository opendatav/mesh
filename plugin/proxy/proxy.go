/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package proxy

import (
	"context"
	"crypto/x509"
	_ "embed"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os/signal"
	"slices"
	"strings"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/go-acme/lego/v4/challenge"
	gokitmetrics "github.com/go-kit/kit/metrics"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/log"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/plugin"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/tool"
	mtypes "github.com/opendatav/mesh/client/golang/types"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/traefik/traefik/v3/cmd/healthcheck"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/metrics"
	"github.com/traefik/traefik/v3/pkg/middlewares/accesslog"
	"github.com/traefik/traefik/v3/pkg/provider/acme"
	"github.com/traefik/traefik/v3/pkg/provider/aggregator"
	"github.com/traefik/traefik/v3/pkg/provider/tailscale"
	"github.com/traefik/traefik/v3/pkg/provider/traefik"
	pxy "github.com/traefik/traefik/v3/pkg/proxy"
	"github.com/traefik/traefik/v3/pkg/proxy/httputil"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
	"github.com/traefik/traefik/v3/pkg/server/router/tcp"
	"github.com/traefik/traefik/v3/pkg/server/service"
	pcp "github.com/traefik/traefik/v3/pkg/tcp"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/tracing"
	"github.com/traefik/traefik/v3/pkg/types"
	"github.com/traefik/traefik/v3/pkg/udp"
	"github.com/traefik/traefik/v3/pkg/version"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	plugin.Provide(proxy)
}

const (
	ProviderName    = "mesh"
	TransportX      = "transport_x"
	TransportY      = "transport_y"
	TransportA      = "transport_a"
	TransportB      = "transport_b"
	TransportC      = "traefik"
	TransportD      = "transport_d"
	Letsencrypt     = "letsencrypt"
	PluginRetry     = "retry"
	PluginBarrier   = "barrier"
	PluginHeader    = "header"
	PluginAuthority = "authority"
	PluginErrors    = "errors"
	PluginRewrite   = "rewrite"
	PluginReplace   = "replace"
	PluginHath      = "hath"
)

var proxy = new(meshProxy)

type meshProxy struct {
	TransportX           string                 `json:"plugin.proxy.transport_x" dft:"0.0.0.0:7305" usage:"Encrypt data frame message address"`
	TransportY           string                 `json:"plugin.proxy.transport_y" dft:"0.0.0.0:7304" usage:"Data frame message address"`
	TransportA           string                 `json:"plugin.proxy.transport_a" dft:"0.0.0.0:443" usage:"User frontend message address"`
	TransportB           string                 `json:"plugin.proxy.transport_b" dft:"0.0.0.0:80" usage:"User frontend message address"`
	TransportC           string                 `json:"plugin.proxy.transport_c" dft:"0.0.0.0:8866" usage:"Internal service address"`
	TransportD           string                 `json:"plugin.proxy.transport_d" dft:"0.0.0.0:8867" usage:"Frontend service address"`
	Home                 string                 `json:"plugin.proxy.home" dft:"${MESH_HOME}/mesh/proxy/" usage:"path to store certification"`
	MaxConcurrencyStream int32                  `json:"plugin.proxy.stream.max" dft:"400" usage:"Max concurrency stream per connection"`
	MaxTransport         int                    `json:"plugin.proxy.transport.max" dft:"400" usage:"Max transports"`
	StatefulEnable       bool                   `json:"plugin.proxy.stateful.enable" dft:"true" usage:"Stateful route enable"`
	RateLimitEnable      bool                   `json:"plugin.proxy.limit.enable" dft:"true" usage:"Rate limit enable"`
	InsecureEnable       bool                   `json:"plugin.proxy.insecure.enable" dft:"true" usage:"Insecure enable"`
	AccessPath           string                 `json:"plugin.proxy.access.path" dft:"${MESH_HOME}/mesh/proxy/logs/access.log" usage:"path to store access log"`
	AccessBackups        int                    `json:"plugin.proxy.access.backups" dft:"120" usage:"Max archive files backups"`
	AccessSize           int                    `json:"plugin.proxy.access.size" dft:"100" usage:"Max archive files size in megabytes"`
	AccessAge            int                    `json:"plugin.proxy.access.age" dft:"28" usage:"Max archive files age in days"`
	Compress             bool                   `json:"plugin.proxy.access.compress" dft:"false" comment:"Compress the accesslog"`
	InsecureSkip         bool                   `json:"plugin.proxy.insecure.skip" dft:"false" comment:"Insecure skip verify"`
	Server               *server.Server         `json:"-"`
	TCPRouters           map[string]*tcp.Router `json:"-"`
	UDPRouters           map[string]udp.Handler `json:"-"`
	RollingWriter        io.WriteCloser         `json:"-"`
}

func (that *meshProxy) Ptt() *plugin.Ptt {
	return &plugin.Ptt{Name: plugin.Proxy, Flags: meshProxy{}, Create: func() plugin.Plugin {
		return that
	}}
}

func (that *meshProxy) Start(ctx context.Context, runtime plugin.Runtime) {
	log.Catch(runtime.Parse(that))
	config := that.Configuration()
	config.SetEffectiveConfiguration()
	err := config.ValidateConfiguration()
	if nil != err {
		log.Error(ctx, "%s", err)
		return
	}
	if err = tool.MakeDir(that.Home); nil != err {
		log.Error(ctx, err.Error())
		return
	}
	// 11G, 120 backups, 100M
	that.RollingWriter = &lumberjack.Logger{
		Filename:   that.AccessPath,
		MaxSize:    that.AccessSize, // megabytes
		MaxBackups: that.AccessBackups,
		MaxAge:     that.AccessAge, //days
		Compress:   that.Compress,  // disabled by default
	}
	accessLog, err := accesslog.NewHandlerWithFormatWriter(config.AccessLog, that.RollingWriter, logger)
	if nil != err {
		log.Error(ctx, "%s", err)
		return
	}
	tls := NewTLSManager(traefiktls.NewManager())
	pb := &proxyBuilder{
		option: &SetupOption{
			SC:       config,
			Provider: meshGraph,
			TCPUpdate: func(tcp map[string]*tcp.Router) {
				that.TCPRouters = tcp
			},
			UDPUpdate: func(udp map[string]udp.Handler) {
				that.UDPRouters = udp
			},
			Setup: func(tcm *pcp.DialerManager) {
			},
		},
	}
	if that.Server, err = pb.build(ctx, accessLog, tls); nil != err {
		log.Error(ctx, err.Error())
		return
	}
	runtime.Submit(func() {
		http.DefaultTransport.(*http.Transport).Proxy = http.ProxyFromEnvironment
		that.startTraefik(ctx, config)
	})
	runtime.StartHook(func() {
		taskId, err := aware.Scheduler.Period(ctx, time.Second*10, &mtypes.Topic{
			Topic: prsim.RoutePeriodRefresh.Topic,
			Code:  prsim.RoutePeriodRefresh.Code,
		})
		if nil != err {
			log.Error(ctx, err.Error())
		} else {
			log.Info(ctx, "Mesh proxy dynamic routers refresh start with %s. ", taskId)
		}
	})
	that.Register(ctx)
}

func (that *meshProxy) Stop(ctx context.Context, runtime plugin.Runtime) {
	if nil != that.Server {
		that.Server.Close()
	}
	if nil != that.RollingWriter {
		log.Catch(that.RollingWriter.Close())
	}
}

func (that *meshProxy) RecursionBreak(ctx context.Context, uri mtypes.URC) bool {
	for _, addr := range uri.Members() {
		if strings.Index(addr, "0.0.0.0") < 0 || strings.Index(addr, "127.0.0.1") < 0 {
			continue
		}
		adds := strings.Split(addr, ":")
		if len(adds) < 2 {
			continue
		}
		if strings.Contains(that.TransportX, fmt.Sprintf(":%s", adds[1])) || strings.Contains(that.TransportY, fmt.Sprintf(":%s", adds[1])) {
			return true
		}
	}
	return false
}

func (that *meshProxy) startTraefik(ctx context.Context, staticConfiguration *static.Configuration) {
	ntx, _ := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	if staticConfiguration.Ping != nil {
		staticConfiguration.Ping.WithContext(ntx)
	}
	that.Server.Start(ntx)
	sent, err := daemon.SdNotify(false, "READY=1")
	if !sent && nil != err {
		log.Error(ntx, "Failed to notify: %v", err)
	}
	t, err := daemon.SdWatchdogEnabled(false)
	if nil != err {
		log.Error(ntx, "Could not enable Watchdog: %v", err)
	} else if t != 0 {
		// Send a ping each half time given
		t /= 2
		log.Info(ntx, "Watchdog activated with timer duration %s", t)
		safe.Go(func() {
			tick := time.Tick(t)
			for range tick {
				resp, errHealthCheck := healthcheck.Do(*staticConfiguration)
				if resp != nil {
					_ = resp.Body.Close()
				}

				if staticConfiguration.Ping == nil || errHealthCheck == nil {
					if ok, _ := daemon.SdNotify(false, "WATCHDOG=1"); !ok {
						log.Error(ntx, "Fail to tick watchdog")
					}
				} else {
					log.Error(ntx, errHealthCheck.Error())
				}
			}
		})
	}
	log.Info(ntx, "Plugin proxy has been started.")
	that.Server.Wait()
	log.Info(ntx, "Plugin proxy has been shutdown")
}

type proxyBuilder struct {
	option *SetupOption
}

func (that *proxyBuilder) build(ctx context.Context, accessLog accesslog.Accesslog, tlsManager TLSManager) (*server.Server, error) {
	providerAggregator := aggregator.NewProviderAggregator(*that.option.SC.Providers)
	// adds internal provider
	if err := providerAggregator.AddProvider(traefik.New(*that.option.SC)); nil != err {
		return nil, cause.Error(err)
	}
	if err := providerAggregator.AddProvider(that.option.Provider); nil != err {
		return nil, cause.Error(err)
	}
	// ACME
	httpChallengeProvider := acme.NewChallengeHTTP()
	tlsChallengeProvider := acme.NewChallengeTLSALPN()
	if err := providerAggregator.AddProvider(tlsChallengeProvider); nil != err {
		return nil, cause.Error(err)
	}
	acmeProviders := that.initACMEProvider(ctx, that.option.SC, providerAggregator, tlsManager.Obtain(ctx), httpChallengeProvider, tlsChallengeProvider)
	// Tailscale
	tsProviders := initTailscaleProviders(ctx, that.option.SC, providerAggregator)
	// Observability
	metricRegistries := that.registerMetricClients(ctx, that.option.SC.Metrics)
	semConvMetricRegistry, err := func() (*metrics.SemConvMetricsRegistry, error) {
		if that.option.SC.Metrics != nil && that.option.SC.Metrics.OTLP != nil {
			semConvMetricRegistry, err := metrics.NewSemConvMetricRegistry(ctx, that.option.SC.Metrics.OTLP)
			return semConvMetricRegistry, cause.Error(err)
		}
		return nil, nil
	}()
	if nil != err {
		return nil, cause.Error(err)
	}
	metricsRegistry := metrics.NewMultiRegistry(metricRegistries)

	tracer, tracerCloser := setupTracing(ctx, that.option.SC.Tracing)
	observabilityMgr := middleware.NewObservabilityMgr(*that.option.SC, metricsRegistry, semConvMetricRegistry, accessLog, tracer, tracerCloser)

	// Entrypoints
	serverEntryPointsTCP, err := server.NewTCPEntryPoints(that.option.SC.EntryPoints, that.option.SC.HostResolver, metricsRegistry)
	if nil != err {
		return nil, cause.Error(err)
	}
	serverEntryPointsUDP, err := server.NewUDPEntryPoints(that.option.SC.EntryPoints)
	if nil != err {
		return nil, cause.Error(err)
	}

	if that.option.SC.API != nil {
		version.DisableDashboardAd = that.option.SC.API.DisableDashboardAd
	}

	// Plugins
	hasPlugins := that.option.SC.Experimental != nil && (that.option.SC.Experimental.Plugins != nil || that.option.SC.Experimental.LocalPlugins != nil)
	if hasPlugins {
		pluginsList := slices.Collect(maps.Keys(that.option.SC.Experimental.Plugins))
		pluginsList = append(pluginsList, slices.Collect(maps.Keys(that.option.SC.Experimental.LocalPlugins))...)
		log.Info(ctx, "Loading plugins: %s", strings.Join(pluginsList, ","))
	}
	pluginBuilder, err := createPluginBuilder(that.option.SC)
	if err != nil && that.option.SC.Experimental != nil && that.option.SC.Experimental.AbortOnPluginFailure {
		return nil, cause.Errorf("plugin: failed to create plugin builder: %s", err)
	}
	if nil != err {
		log.Error(ctx, "Plugins are disabled because an error has occurred. %s", err.Error())
	}
	// Providers plugins
	for name, conf := range that.option.SC.Providers.Plugin {
		if nil == pluginBuilder {
			break
		}
		p, err := pluginBuilder.BuildProvider(name, conf)
		if nil != err {
			return nil, cause.Errorf("plugin: failed to build provider: %w", err)
		}
		if err = providerAggregator.AddProvider(p); nil != err {
			return nil, cause.Errorf("plugin: failed to add provider: %w", err)
		}
	}

	// Service manager factory
	var spiffeX509Source *workloadapi.X509Source
	if that.option.SC.Spiffe != nil && that.option.SC.Spiffe.WorkloadAPIAddr != "" {
		log.Info(ctx, "Waiting on SPIFFE SVID delivery, workloadAPIAddr is %s", that.option.SC.Spiffe.WorkloadAPIAddr)
		spiffeX509Source, err = workloadapi.NewX509Source(ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(that.option.SC.Spiffe.WorkloadAPIAddr)))
		if nil != err {
			return nil, cause.Errorf("unable to create SPIFFE x509 source: %w", err)
		}
		log.Info(ctx, "Successfully obtained SPIFFE SVID.")
	}

	transportManager := service.NewTransportManager(spiffeX509Source)
	var proxyBuilder service.ProxyBuilder = httputil.NewProxyBuilder(transportManager, semConvMetricRegistry)
	if that.option.SC.Experimental != nil && that.option.SC.Experimental.FastProxy != nil {
		proxyBuilder = pxy.NewSmartBuilder(transportManager, proxyBuilder, *that.option.SC.Experimental.FastProxy)
	}

	dialerManager := pcp.NewDialerManager(spiffeX509Source)
	if nil != that.option.Setup {
		that.option.Setup(dialerManager)
	}
	acmeHTTPHandler := that.getHTTPChallengeHandler(acmeProviders, httpChallengeProvider)

	routinesPool := safe.NewPool(ctx)

	managerFactory := service.NewManagerFactory(*that.option.SC, routinesPool, observabilityMgr, transportManager, proxyBuilder, acmeHTTPHandler)
	// Router factory
	routerFactory := server.NewRouterFactory(*that.option.SC, managerFactory, tlsManager.Obtain(ctx), observabilityMgr, pluginBuilder, dialerManager)
	// Watcher
	watcher := server.NewConfigurationWatcher(routinesPool, providerAggregator, that.getDefaultsEntrypoints(ctx, that.option.SC), "internal")
	// TLS
	watcher.AddListener(func(conf dynamic.Configuration) {
		tlsManager.UpdateConfigs(macro.Context(), conf.TLS.Stores, conf.TLS.Options, conf.TLS.Certificates)
		gauge := metricsRegistry.TLSCertsNotAfterTimestampGauge()
		for _, certificate := range tlsManager.Obtain(ctx).GetServerCertificates() {
			that.appendCertMetric(gauge, certificate)
		}
	})
	// Metrics
	watcher.AddListener(func(_ dynamic.Configuration) {
		metricsRegistry.ConfigReloadsCounter().Add(1)
		metricsRegistry.LastConfigReloadSuccessGauge().Set(float64(time.Now().Unix()))
	})
	// Server Transports
	watcher.AddListener(func(conf dynamic.Configuration) {
		log.Info(ctx, "proxy configuration updating")
		transportManager.Update(conf.HTTP.ServersTransports)
		proxyBuilder.Update(conf.HTTP.ServersTransports)
		dialerManager.Update(conf.TCP.ServersTransports)
	})
	// Switch router
	watcher.AddListener(that.switchRouter(routerFactory, serverEntryPointsTCP, serverEntryPointsUDP))
	// Metrics
	if metricsRegistry.IsEpEnabled() || metricsRegistry.IsRouterEnabled() || metricsRegistry.IsSvcEnabled() {
		var eps []string
		for key := range serverEntryPointsTCP {
			eps = append(eps, key)
		}
		watcher.AddListener(func(conf dynamic.Configuration) {
			metrics.OnConfigurationUpdate(conf, eps)
		})
	}
	// TLS challenge
	watcher.AddListener(tlsChallengeProvider.ListenConfiguration)
	// Certificate Resolvers
	resolverNames := map[string]struct{}{}
	// ACME
	for _, p := range acmeProviders {
		resolverNames[p.ResolverName] = struct{}{}
		watcher.AddListener(p.ListenConfiguration)
	}
	// Tailscale
	for _, p := range tsProviders {
		resolverNames[p.ResolverName] = struct{}{}
		watcher.AddListener(p.HandleConfigUpdate)
	}
	// Certificate resolver logs
	watcher.AddListener(func(config dynamic.Configuration) {
		for rtName, rt := range config.HTTP.Routers {
			if rt.TLS == nil || rt.TLS.CertResolver == "" {
				continue
			}
			if _, ok := resolverNames[rt.TLS.CertResolver]; !ok {
				log.Warn(ctx, "Router %s uses a non-existent certificate resolver: %s", rtName, rt.TLS.CertResolver)
			}
		}
	})
	return server.NewServer(routinesPool, serverEntryPointsTCP, serverEntryPointsUDP, watcher, observabilityMgr), nil
}

func (that *proxyBuilder) getHTTPChallengeHandler(acmeProviders []*acme.Provider, httpChallengeProvider http.Handler) http.Handler {
	var acmeHTTPHandler http.Handler
	for _, p := range acmeProviders {
		if p != nil && p.HTTPChallenge != nil {
			acmeHTTPHandler = httpChallengeProvider
			break
		}
	}
	return acmeHTTPHandler
}

func (that *proxyBuilder) getDefaultsEntrypoints(ctx context.Context, staticConfiguration *static.Configuration) []string {
	var defaultEntryPoints []string

	// Determines if at least one EntryPoint is configured to be used by default.
	var hasDefinedDefaults bool
	for _, ep := range staticConfiguration.EntryPoints {
		if ep.AsDefault {
			hasDefinedDefaults = true
			break
		}
	}

	for name, cfg := range staticConfiguration.EntryPoints {
		// By default, all entrypoints are considered.
		// If at least one is flagged, then only flagged entrypoints are included.
		if hasDefinedDefaults && !cfg.AsDefault {
			continue
		}

		protocol, err := cfg.GetProtocol()
		if nil != err {
			// Should never happen because Traefik should not start if protocol is invalid.
			log.Error(ctx, "Invalid protocol: %v", err)
		}
		if protocol != "udp" && name != static.DefaultInternalEntryPointName {
			defaultEntryPoints = append(defaultEntryPoints, name)
		}
	}
	slices.Sort(defaultEntryPoints)
	return defaultEntryPoints
}

func (that *proxyBuilder) switchRouter(routerFactory *server.RouterFactory, serverEntryPointsTCP server.TCPEntryPoints, serverEntryPointsUDP server.UDPEntryPoints) func(conf dynamic.Configuration) {
	return func(conf dynamic.Configuration) {
		rtConf := runtime.NewConfig(conf)
		routers, udpRouters := routerFactory.CreateRouters(rtConf)
		serverEntryPointsTCP.Switch(routers)
		serverEntryPointsUDP.Switch(udpRouters)
		if x := that.option.TCPUpdate; nil != x {
			x(routers)
		}
		if x := that.option.UDPUpdate; nil != x {
			x(udpRouters)
		}
	}
}

// initACMEProvider creates and registers acme.Provider instances corresponding to the configured ACME certificate resolvers.
func (that *proxyBuilder) initACMEProvider(ctx context.Context, c *static.Configuration, providerAggregator *aggregator.ProviderAggregator, tlsManager *traefiktls.Manager, httpChallengeProvider, tlsChallengeProvider challenge.Provider) []*acme.Provider {
	localStores := map[string]acme.Store{}
	var resolvers []*acme.Provider
	for name, resolver := range c.CertificatesResolvers {
		if resolver.ACME == nil {
			continue
		}
		if localStores[resolver.ACME.Storage] == nil {
			localStores[resolver.ACME.Storage] = acme.NewLocalStore(resolver.ACME.Storage)
		}
		p := &acme.Provider{
			Configuration:         resolver.ACME,
			Store:                 localStores[resolver.ACME.Storage],
			ResolverName:          name,
			HTTPChallengeProvider: httpChallengeProvider,
			TLSChallengeProvider:  tlsChallengeProvider,
		}
		if err := providerAggregator.AddProvider(p); nil != err {
			log.Error(ctx, "The ACME resolver %q is skipped from the resolvers list because: %v", name, err)
			continue
		}
		p.SetTLSManager(tlsManager)
		p.SetConfigListenerChan(make(chan dynamic.Configuration))
		resolvers = append(resolvers, p)
	}
	return resolvers
}

// initTailscaleProviders creates and registers tailscale.Provider instances corresponding to the configured Tailscale certificate resolvers.
func initTailscaleProviders(ctx context.Context, cfg *static.Configuration, providerAggregator *aggregator.ProviderAggregator) []*tailscale.Provider {
	var providers []*tailscale.Provider
	for name, resolver := range cfg.CertificatesResolvers {
		if resolver.Tailscale == nil {
			continue
		}

		tsProvider := &tailscale.Provider{ResolverName: name}

		if err := providerAggregator.AddProvider(tsProvider); nil != err {
			log.Error(ctx, "Unable to create Tailscale provider %s: %v", name, err)
			continue
		}

		providers = append(providers, tsProvider)
	}

	return providers
}

func (that *proxyBuilder) registerMetricClients(ctx context.Context, metricsConfig *types.Metrics) []metrics.Registry {
	if metricsConfig == nil {
		return nil
	}
	var registries []metrics.Registry
	if metricsConfig.Prometheus != nil {
		prometheusRegister := metrics.RegisterPrometheus(ctx, metricsConfig.Prometheus)
		if prometheusRegister != nil {
			registries = append(registries, prometheusRegister)
			log.Debug(ctx, "Configured Prometheus metrics")
		}
	}
	if metricsConfig.Datadog != nil {
		registries = append(registries, metrics.RegisterDatadog(ctx, metricsConfig.Datadog))
		log.Debug(ctx, "Configured Datadog metrics: pushing to %s once every %s",
			metricsConfig.Datadog.Address, metricsConfig.Datadog.PushInterval)
	}
	if metricsConfig.StatsD != nil {
		registries = append(registries, metrics.RegisterStatsd(ctx, metricsConfig.StatsD))
		log.Debug(ctx, "Configured StatsD metrics: pushing to %s once every %s",
			metricsConfig.StatsD.Address, metricsConfig.StatsD.PushInterval)
	}
	if metricsConfig.InfluxDB2 != nil {
		influxDB2Register := metrics.RegisterInfluxDB2(ctx, metricsConfig.InfluxDB2)
		if influxDB2Register != nil {
			registries = append(registries, influxDB2Register)
			log.Debug(ctx, "Configured InfluxDB v2 metrics: pushing to %s (%s org/%s bucket) once every %s",
				metricsConfig.InfluxDB2.Address, metricsConfig.InfluxDB2.Org, metricsConfig.InfluxDB2.Bucket, metricsConfig.InfluxDB2.PushInterval)
		}
	}
	if metricsConfig.OTLP != nil {
		openTelemetryRegistry := metrics.RegisterOpenTelemetry(ctx, metricsConfig.OTLP)
		if openTelemetryRegistry != nil {
			registries = append(registries, openTelemetryRegistry)
			log.Debug(ctx, "Configured OpenTelemetry metrics: pushing to %s once every %s",
				metricsConfig.OTLP.ServiceName, metricsConfig.OTLP.PushInterval.String())
		}
	}
	return registries
}

func (that *proxyBuilder) appendCertMetric(gauge gokitmetrics.Gauge, certificate *x509.Certificate) {
	slices.Sort(certificate.DNSNames)
	labels := []string{
		"cn", certificate.Subject.CommonName,
		"serial", certificate.SerialNumber.String(),
		"sans", strings.Join(certificate.DNSNames, ","),
	}
	notAfter := float64(certificate.NotAfter.Unix())
	gauge.With(labels...).Set(notAfter)
}

func setupTracing(ctx context.Context, conf *static.Tracing) (*tracing.Tracer, io.Closer) {
	if nil == conf {
		return nil, nil
	}
	tracer, closer, err := tracing.NewTracing(conf)
	if err != nil {
		log.Warn(ctx, "Unable to create tracer")
		return nil, nil
	}

	return tracer, closer
}
