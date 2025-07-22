/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package metabase

import (
	"context"
	"github.com/opendatav/mesh/client/golang/plugin"
)

func init() {
	plugin.Provide(metabase)
}

const Name = "metabase"

var metabase = new(meta)

type meta struct {
}

func (that *meta) Ptt() *plugin.Ptt {
	return &plugin.Ptt{Name: Name, Flags: meta{}, Create: func() plugin.Plugin {
		return that
	}}
}

func (that *meta) Start(ctx context.Context, runtime plugin.Runtime) {

}

func (that *meta) Stop(ctx context.Context, runtime plugin.Runtime) {

}
