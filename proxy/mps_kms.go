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
	"github.com/opendatav/mesh/client/golang/types"
	prsimmps "github.com/opendatav/mesh/plugin/prsim"
)

func init() {
	var service = &meshPRSIKMSMPS{mps: new(prsimmps.PRSIKMS)}
	var _ prsim.KMS = service
	macro.Provide((*prsim.KMS)(nil), service)

	var serviceReset = &meshPRSIKMSResetMPS{service: service}
	var _ mpc.Invoker = serviceReset
	var _ macro.Caller = serviceReset
	macro.Provide(mpc.IInvoker, serviceReset)
	macro.Provide(macro.ICaller, serviceReset)

	var serviceEnviron = &meshPRSIKMSEnvironMPS{service: service}
	var _ mpc.Invoker = serviceEnviron
	var _ macro.Caller = serviceEnviron
	macro.Provide(mpc.IInvoker, serviceEnviron)
	macro.Provide(macro.ICaller, serviceEnviron)

	var serviceList = &meshPRSIKMSListMPS{service: service}
	var _ mpc.Invoker = serviceList
	var _ macro.Caller = serviceList
	macro.Provide(mpc.IInvoker, serviceList)
	macro.Provide(macro.ICaller, serviceList)

	var serviceApplyRoot = &meshPRSIKMSApplyRootMPS{service: service}
	var _ mpc.Invoker = serviceApplyRoot
	var _ macro.Caller = serviceApplyRoot
	macro.Provide(mpc.IInvoker, serviceApplyRoot)
	macro.Provide(macro.ICaller, serviceApplyRoot)

	var serviceApplyIssue = &meshPRSIKMSApplyIssueMPS{service: service}
	var _ mpc.Invoker = serviceApplyIssue
	var _ macro.Caller = serviceApplyIssue
	macro.Provide(mpc.IInvoker, serviceApplyIssue)
	macro.Provide(macro.ICaller, serviceApplyIssue)

}

// meshPRSIKMSMPS is an implementation of prsim.PRSIKMS
type meshPRSIKMSMPS struct {
	mps *prsimmps.PRSIKMS
}

func (that *meshPRSIKMSMPS) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshSPI}
}

func (that *meshPRSIKMSMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: macro.MeshMPS}
}

// Reset will override the keystore environ.
// @MPI("kms.store.reset")
func (that *meshPRSIKMSMPS) Reset(ctx context.Context, env *types.Environ) error {
	return that.mps.Reset(ctx, env)
}

// Environ will return the keystore environ.
// @MPI("kms.store.environ")
func (that *meshPRSIKMSMPS) Environ(ctx context.Context) (*types.Environ, error) {
	return that.mps.Environ(ctx)
}

// List will return the keystore environ.
// @MPI("kms.crt.store.list")
func (that *meshPRSIKMSMPS) List(ctx context.Context, cno string) ([]*types.Keys, error) {
	return that.mps.List(ctx, cno)
}

// ApplyRoot will apply the root certification.
// @MPI("kms.crt.apply.root")
func (that *meshPRSIKMSMPS) ApplyRoot(ctx context.Context, csr *types.KeyCsr) ([]*types.Keys, error) {
	return that.mps.ApplyRoot(ctx, csr)
}

// ApplyIssue will apply the common certification.
// @MPI("kms.crt.apply.issue")
func (that *meshPRSIKMSMPS) ApplyIssue(ctx context.Context, csr *types.KeyCsr) ([]*types.Keys, error) {
	return that.mps.ApplyIssue(ctx, csr)
}

type meshPRSIKMSResetMPS struct {
	service prsim.KMS
}

func (that *meshPRSIKMSResetMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("kms.store.reset", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSIKMSResetMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("kms.store.reset", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSIKMSResetMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshKMSResetParameters)
	parameters.SetArguments(ctx, args...)
	return nil, that.service.Reset(ctx, parameters.Env)
}

func (that *meshPRSIKMSResetMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshKMSResetParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return nil, that.service.Reset(ctx, parameters.Env)
}

type meshPRSIKMSEnvironMPS struct {
	service prsim.KMS
}

func (that *meshPRSIKMSEnvironMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("kms.store.environ", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSIKMSEnvironMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("kms.store.environ", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSIKMSEnvironMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	return that.service.Environ(ctx)
}

func (that *meshPRSIKMSEnvironMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	return that.service.Environ(ctx)
}

type meshPRSIKMSListMPS struct {
	service prsim.KMS
}

func (that *meshPRSIKMSListMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("kms.crt.store.list", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSIKMSListMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("kms.crt.store.list", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSIKMSListMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshKMSListParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.List(ctx, parameters.Cno)
}

func (that *meshPRSIKMSListMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshKMSListParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.List(ctx, parameters.Cno)
}

type meshPRSIKMSApplyRootMPS struct {
	service prsim.KMS
}

func (that *meshPRSIKMSApplyRootMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("kms.crt.apply.root", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSIKMSApplyRootMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("kms.crt.apply.root", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSIKMSApplyRootMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshKMSApplyRootParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.ApplyRoot(ctx, parameters.Csr)
}

func (that *meshPRSIKMSApplyRootMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshKMSApplyRootParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.ApplyRoot(ctx, parameters.Csr)
}

type meshPRSIKMSApplyIssueMPS struct {
	service prsim.KMS
}

func (that *meshPRSIKMSApplyIssueMPS) Att() *macro.Att {
	return &macro.Att{Name: strings.ReplaceAll("kms.crt.apply.issue", "${mesh.name}", tool.Name.Get())}
}
func (that *meshPRSIKMSApplyIssueMPS) Stt() *macro.Stt {
	return &macro.Stt{Name: strings.ReplaceAll("kms.crt.apply.issue", "${mesh.name}", tool.Name.Get())}
}

func (that *meshPRSIKMSApplyIssueMPS) Call(ctx context.Context, proxy interface{}, method macro.Inspector, args ...interface{}) (interface{}, error) {
	parameters := new(proxympx.MeshKMSApplyIssueParameters)
	parameters.SetArguments(ctx, args...)
	return that.service.ApplyIssue(ctx, parameters.Csr)
}

func (that *meshPRSIKMSApplyIssueMPS) Invoke(ctx context.Context, invocation mpc.Invocation) (interface{}, error) {
	parameters, ok := invocation.GetParameters().(*proxympx.MeshKMSApplyIssueParameters)
	if !ok {
		return nil, cause.CompatibleError("Service %s parameters can't compatible. ", mpc.ContextWith(ctx).GetUrn())
	}
	return that.service.ApplyIssue(ctx, parameters.Csr)
}
