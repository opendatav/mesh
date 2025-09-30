/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package proxy

import (
	"context"
	"sync"

	"github.com/opendatav/mesh/client/golang/tool"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/provider"
	stc "github.com/traefik/traefik/v3/pkg/server/router/tcp"
	"github.com/traefik/traefik/v3/pkg/tcp"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/udp"
)

type SetupOption struct {
	SC        *static.Configuration
	Provider  provider.Provider
	TCPUpdate func(tcp map[string]*stc.Router)
	UDPUpdate func(udp map[string]udp.Handler)
	Setup     func(tcm *tcp.DialerManager)
	ShareTLS  bool
}

func init() {
	var _ TLSManager = new(sharedTLSManager)
}

var DefaultTLSOptions = tls.Options{
	// ensure http2 enabled
	ALPNProtocols: []string{"h2", "http/1.1"},
	MinVersion:    "VersionTLS12",
	CipherSuites:  tls.DefaultTLSOptions.CipherSuites,
}

func NewUnsharedTLSManager(tm *tls.Manager) TLSManager {
	return &unsharedTLSManager{
		tls: tm,
	}
}

func NewTLSManager(tm *tls.Manager) TLSManager {
	return &sharedTLSManager{
		lock:    sync.RWMutex{},
		tls:     tm,
		stores:  map[string]tls.Store{},
		configs: map[string]tls.Options{},
		certs:   map[string]*tls.CertAndStores{},
	}
}

type TLSManager interface {
	Obtain(ctx context.Context) *tls.Manager
	UpdateConfigs(ctx context.Context, stores map[string]tls.Store, configs map[string]tls.Options, certs []*tls.CertAndStores)
}

type sharedTLSManager struct {
	lock    sync.RWMutex
	tls     *tls.Manager
	stores  map[string]tls.Store
	configs map[string]tls.Options
	certs   map[string]*tls.CertAndStores
}

func (that *sharedTLSManager) Obtain(ctx context.Context) *tls.Manager {
	return that.tls
}

func (that *sharedTLSManager) UpdateConfigs(ctx context.Context, stores map[string]tls.Store, configs map[string]tls.Options, certs []*tls.CertAndStores) {
	that.lock.Lock()
	defer that.lock.Unlock()
	for k, v := range stores {
		that.stores[k] = v
	}
	for k, v := range configs {
		that.configs[k] = v
	}
	for _, cert := range certs {
		that.certs[cert.CertFile.String()] = cert
	}
	xs := map[string]tls.Store{}
	for k, v := range that.stores {
		xs[k] = v
	}
	xc := map[string]tls.Options{}
	for k, v := range that.configs {
		xc[k] = v
	}
	that.tls.UpdateConfigs(ctx, xs, xc, tool.Values(that.certs))
}

type unsharedTLSManager struct {
	tls *tls.Manager
}

func (that *unsharedTLSManager) Obtain(ctx context.Context) *tls.Manager {
	return that.tls
}

func (that *unsharedTLSManager) UpdateConfigs(ctx context.Context, stores map[string]tls.Store, configs map[string]tls.Options, certs []*tls.CertAndStores) {
	that.tls.UpdateConfigs(ctx, stores, configs, certs)
}
