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

	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/mpc"
	proxympx "github.com/opendatav/mesh/client/golang/proxy"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/tool"
	_ "github.com/opendatav/mesh/plugin/metabase"
	prsimmps "github.com/opendatav/mesh/plugin/prsim"
)

func init() {
	var service = &meshPRSISequenceMPS{mps: new(prsimmps.PRSISequence)}
	var _ prsim.Sequence = service
	macro.Provide((*prsim.Sequence)(nil), service)

	var serviceNext = &meshPRSISequenceNextMPS{service: service}
	var _ mpc.Invoker = serviceNext
	var _ macro.Caller = serviceNext
	macro.Provide(mpc.IInvoker, serviceNext)
	macro.Provide(macro.ICaller, serviceNext)

	var serviceSection = &meshPRSISequenceSectionMPS{service: service}
	var _ mpc.Invoker = serviceSection
	var _ macro.Caller = serviceSection
	macro.Provide(mpc.IInvoker, serviceSection)
	macro.Provide(macro.ICaller, serviceSection)

}

// meshPRSISequenceMPS is an implementation of prsim.PRSISequence
type meshPRSISequenceMPS struct {
	mps *prsimmps.PRSISequence
}

func (that *meshPRSISequenceMPS) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshSPI}
}

func (that *meshPRSISequenceMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: macro.MeshMPS}
}

// Next
// @MPI("mesh.sequence.next")
func (that *meshPRSISequenceMPS) Next(ctx context.Context, kind string, length int) (string, error) {
	return that.mps.Next(ctx, kind, length)
}

// Section
// @MPI("mesh.sequence.section")
func (that *meshPRSISequenceMPS) Section(ctx context.Context, kind string, size int, length int) ([]string, error) {
	return that.mps.Section(ctx, kind, size, length)
}

type meshPRSISequenceNextMPS struct {
	service prsim.Sequence
}

func (that *meshPRSISequenceNextMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("mesh.sequence.next", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSISequenceNextMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("mesh.sequence.next", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSISequenceNextMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshSequenceNextParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.Next(ctx, parameters.Kind, parameters.Length)
}

func (that *meshPRSISequenceNextMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshSequenceNextParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.Next(ctx, parameters.Kind, parameters.Length)
}

type meshPRSISequenceSectionMPS struct {
	service prsim.Sequence
}

func (that *meshPRSISequenceSectionMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("mesh.sequence.section", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSISequenceSectionMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("mesh.sequence.section", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSISequenceSectionMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshSequenceSectionParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.Section(ctx, parameters.Kind, parameters.Size, parameters.Length)
}

func (that *meshPRSISequenceSectionMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshSequenceSectionParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.Section(ctx, parameters.Kind, parameters.Size, parameters.Length)
}
