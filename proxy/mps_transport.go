/*
* Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
* TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
*
*
 */

// Code generated by mesh; DO NOT EDIT.

package proxy

import (
	"context"
	"strings"

	"github.com/be-io/mesh/client/golang/cause"
	"github.com/be-io/mesh/client/golang/macro"
	"github.com/be-io/mesh/client/golang/mpc"
	proxympx "github.com/be-io/mesh/client/golang/proxy"
	"github.com/be-io/mesh/client/golang/prsim"
	"github.com/be-io/mesh/client/golang/tool"
	"github.com/be-io/mesh/client/golang/types"
	prsimmps "github.com/be-io/mesh/plugin/prsim"
)

func init() {
	var service = &meshPRSITransportMPS{mps: new(prsimmps.PRSITransport)}
	var _ prsim.Transport = service
	macro.Provide((*prsim.Transport)(nil), service)

	var serviceOpen = &meshPRSITransportOpenMPS{service: service}
	var _ mpc.Invoker = serviceOpen
	var _ macro.Caller = serviceOpen
	macro.Provide(mpc.IInvoker, serviceOpen)
	macro.Provide(macro.ICaller, serviceOpen)

	var serviceClose = &meshPRSITransportCloseMPS{service: service}
	var _ mpc.Invoker = serviceClose
	var _ macro.Caller = serviceClose
	macro.Provide(mpc.IInvoker, serviceClose)
	macro.Provide(macro.ICaller, serviceClose)

	var serviceRoundtrip = &meshPRSITransportRoundtripMPS{service: service}
	var _ mpc.Invoker = serviceRoundtrip
	var _ macro.Caller = serviceRoundtrip
	macro.Provide(mpc.IInvoker, serviceRoundtrip)
	macro.Provide(macro.ICaller, serviceRoundtrip)

}

// meshPRSITransportMPS is an implementation of prsim.PRSITransport
type meshPRSITransportMPS struct {
	mps *prsimmps.PRSITransport
}

func (that *meshPRSITransportMPS) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshSPI}
}

func (that *meshPRSITransportMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: macro.MeshMPS}
}

// Open a channel session.
// @MPI("mesh.chan.open")
func (that *meshPRSITransportMPS) Open(ctx context.Context, sessionId string, metadata map[string]string) (prsim.Session, error) {
	return that.mps.Open(ctx, sessionId, metadata)
}

// Close the channel.
// @MPI("mesh.chan.close")
func (that *meshPRSITransportMPS) Close(ctx context.Context, timeout types.Duration) error {
	return that.mps.Close(ctx, timeout)
}

// Roundtrip with the channel.
// @MPI("mesh.chan.roundtrip")
func (that *meshPRSITransportMPS) Roundtrip(ctx context.Context, payload []byte, metadata map[string]string) ([]byte, error) {
	return that.mps.Roundtrip(ctx, payload, metadata)
}

type meshPRSITransportOpenMPS struct {
	service prsim.Transport
}

func (that *meshPRSITransportOpenMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("mesh.chan.open", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSITransportOpenMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("mesh.chan.open", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSITransportOpenMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshTransportOpenParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.Open(ctx, parameters.SessionId, parameters.Metadata)
}

func (that *meshPRSITransportOpenMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshTransportOpenParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.Open(ctx, parameters.SessionId, parameters.Metadata)
}

type meshPRSITransportCloseMPS struct {
	service prsim.Transport
}

func (that *meshPRSITransportCloseMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("mesh.chan.close", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSITransportCloseMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("mesh.chan.close", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSITransportCloseMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshTransportCloseParameters)
	parameters.SetArguments(ctx, args...)
	return nil, that.service.Close(ctx, parameters.Timeout)
}

func (that *meshPRSITransportCloseMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshTransportCloseParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return nil, that.service.Close(ctx, parameters.Timeout)
}

type meshPRSITransportRoundtripMPS struct {
	service prsim.Transport
}

func (that *meshPRSITransportRoundtripMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("mesh.chan.roundtrip", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSITransportRoundtripMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("mesh.chan.roundtrip", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSITransportRoundtripMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshTransportRoundtripParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.Roundtrip(ctx, parameters.Payload, parameters.Metadata)
}

func (that *meshPRSITransportRoundtripMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshTransportRoundtripParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.Roundtrip(ctx, parameters.Payload, parameters.Metadata)
}
