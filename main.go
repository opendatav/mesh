/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package main

import (
	"git.firmer.tech/firmer/mesh/client/golang/log"
	"git.firmer.tech/firmer/mesh/client/golang/mesh"
	"github.com/opendatav/mesh/client/golang/mpc"
	"github.com/opendatav/mesh/client/golang/plugin"
	//
	_ "git.firmer.tech/firmer/mesh/client/golang/boost"
	mm "git.firmer.tech/firmer/mesh/client/golang/macro"
	_ "git.firmer.tech/firmer/mesh/client/golang/orm/dao"
	_ "git.firmer.tech/firmer/mesh/client/golang/proxy"
	_ "git.firmer.tech/firmer/mesh/client/golang/system"
	//
	_ "github.com/opendatav/mesh/cmd"
	_ "github.com/opendatav/mesh/plugin/doc"
	_ "github.com/opendatav/mesh/plugin/webapp"
)

//go:generate go run client/golang/dyn/mpc.go -i plugin -e plugin/raft -m github.com/opendatav/mesh

func main() {
	if err := mm.Active(); nil != err {
		panic(err)
	}
	ctx := mesh.Context()
	hooks := mm.Load[mm.RuntimeHook]().List()
	for _, hook := range hooks {
		rh, ok := hook.(mm.RuntimeHook)
		if !ok {
			log.Warnf(ctx, "Hook %s cant be cast to RuntimeHook", hook.Att().Name)
			continue
		}
		if err := rh.Start(ctx, mesh.GoDescriptor); nil != err {
			log.Fatalf(ctx, "Hook %s exec, %s", hook.Att().Name, err)
		}
	}
	container := plugin.LoadC("mesh")
	container.Start(mpc.Context())
}
