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

	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/mpc"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/types"
)

func init() {
	var _ macro.Interface = new(meshDevopsMPI)
	macro.Provide((*prsim.Devops)(nil), &meshDevopsMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"Distribute": {
				DeclaredKind: (*prsim.Devops)(nil),
				TName:        "prsim.Devops",
				Name:         "Distribute",
				Intype:       func() macro.Parameters { var parameters MeshDevopsDistributeParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshDevopsDistributeReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshDevopsDistributeParameters) },
				Outbound:     func() macro.Returns { return new(MeshDevopsDistributeReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.devops.distribute"},
				},
			},
			"Transform": {
				DeclaredKind: (*prsim.Devops)(nil),
				TName:        "prsim.Devops",
				Name:         "Transform",
				Intype:       func() macro.Parameters { var parameters MeshDevopsTransformParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshDevopsTransformReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshDevopsTransformParameters) },
				Outbound:     func() macro.Returns { return new(MeshDevopsTransformReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.devops.transform"},
				},
			},
		},
	})
}

// meshDevopsMPI is an implementation of Devops
type meshDevopsMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshDevopsMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshDevopsMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshDevopsMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// Distribute the service sdk.
// @MPI("mesh.devops.distribute")
func (that *meshDevopsMPI) Distribute(ctx context.Context, option *types.DistributeOption) (string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Distribute"], option)
	if nil != err {
		x := new(MeshDevopsDistributeReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshDevopsDistributeReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshDevopsDistributeReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Transform standard schema to another schema.
// @MPI("mesh.devops.transform")
func (that *meshDevopsMPI) Transform(ctx context.Context, option *types.TransformOption) (string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Transform"], option)
	if nil != err {
		x := new(MeshDevopsTransformReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshDevopsTransformReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshDevopsTransformReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

type MeshDevopsDistributeParameters struct {
	Attachments map[string]string       `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Option      *types.DistributeOption `index:"0" json:"option" xml:"option" yaml:"option"`
}

func (that *MeshDevopsDistributeParameters) GetKind() interface{} {
	return new(MeshDevopsDistributeParameters)
}

func (that *MeshDevopsDistributeParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Option)
	return arguments
}

func (that *MeshDevopsDistributeParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Option = arguments[0].(*types.DistributeOption)
		}
	}
}

func (that *MeshDevopsDistributeParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshDevopsDistributeParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshDevopsDistributeReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content string       `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshDevopsDistributeReturns) GetCode() string {
	return that.Code
}

func (that *MeshDevopsDistributeReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshDevopsDistributeReturns) GetMessage() string {
	return that.Message
}

func (that *MeshDevopsDistributeReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshDevopsDistributeReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshDevopsDistributeReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshDevopsDistributeReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshDevopsDistributeReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(string)
		}
	}
}

type MeshDevopsTransformParameters struct {
	Attachments map[string]string      `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Option      *types.TransformOption `index:"0" json:"option" xml:"option" yaml:"option"`
}

func (that *MeshDevopsTransformParameters) GetKind() interface{} {
	return new(MeshDevopsTransformParameters)
}

func (that *MeshDevopsTransformParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Option)
	return arguments
}

func (that *MeshDevopsTransformParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Option = arguments[0].(*types.TransformOption)
		}
	}
}

func (that *MeshDevopsTransformParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshDevopsTransformParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshDevopsTransformReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content string       `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshDevopsTransformReturns) GetCode() string {
	return that.Code
}

func (that *MeshDevopsTransformReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshDevopsTransformReturns) GetMessage() string {
	return that.Message
}

func (that *MeshDevopsTransformReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshDevopsTransformReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshDevopsTransformReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshDevopsTransformReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshDevopsTransformReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(string)
		}
	}
}
