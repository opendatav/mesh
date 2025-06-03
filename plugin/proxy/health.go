/*
 * Copyright (c) 2000, 2099, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package proxy

import (
	"context"
	"fmt"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/log"
	"github.com/opendatav/mesh/client/golang/mpc"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
	"net/http"
	"strings"
)

func init() {
	var _ http.Handler = new(health)
	middleware.Provide(&healthMiddleware{})
}

type health struct {
	name string
	next http.Handler
}

func (that *health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/f5/checkHealth") {
		if _, err := w.Write([]byte("@the@thealth@is@good@")); nil != err {
			log.Error(mpc.HTTPTracerContext(r), "Health response with %s", err)
		}
		return
	}
	if r.URL.Path == "/stats" {
		if _, err := w.Write([]byte(cause.Success.Code)); nil != err {
			log.Error(mpc.HTTPTracerContext(r), "Health response with %s", err)
		}
		return
	}
	that.next.ServeHTTP(w, r)
}

type healthMiddleware struct {
}

func (that *healthMiddleware) Name() string {
	return fmt.Sprintf("health@%s", ProviderName)
}

func (that *healthMiddleware) Priority() int {
	return 0
}

func (that *healthMiddleware) Scope() int {
	return 0
}

func (that *healthMiddleware) New(ctx context.Context, next http.Handler, name string) (http.Handler, error) {
	return &health{next: next, name: name}, nil
}
