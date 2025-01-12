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
)

func init() {
	var _ macro.Interface = new(meshPreProxyMPI)
	macro.Provide((*prsim.PreProxy)(nil), &meshPreProxyMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"ProxyInvokeGlite2Omega": {
				DeclaredKind: (*prsim.PreProxy)(nil),
				TName:        "prsim.PreProxy",
				Name:         "ProxyInvokeGlite2Omega",
				Intype: func() macro.Parameters {
					var parameters MeshPreProxyProxyInvokeGlite2OmegaParameters
					return &parameters
				},
				Retype:   func() macro.Returns { var returns MeshPreProxyProxyInvokeGlite2OmegaReturns; return &returns },
				Inbound:  func() macro.Parameters { return new(MeshPreProxyProxyInvokeGlite2OmegaParameters) },
				Outbound: func() macro.Returns { return new(MeshPreProxyProxyInvokeGlite2OmegaReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.glite2omega.api"},
				},
			},
			"ProxyInvokeOmega2Glite": {
				DeclaredKind: (*prsim.PreProxy)(nil),
				TName:        "prsim.PreProxy",
				Name:         "ProxyInvokeOmega2Glite",
				Intype: func() macro.Parameters {
					var parameters MeshPreProxyProxyInvokeOmega2GliteParameters
					return &parameters
				},
				Retype:   func() macro.Returns { var returns MeshPreProxyProxyInvokeOmega2GliteReturns; return &returns },
				Inbound:  func() macro.Parameters { return new(MeshPreProxyProxyInvokeOmega2GliteParameters) },
				Outbound: func() macro.Returns { return new(MeshPreProxyProxyInvokeOmega2GliteReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.omega2glite.api"},
				},
			},
			"InvokeGliteConnectivity": {
				DeclaredKind: (*prsim.PreProxy)(nil),
				TName:        "prsim.PreProxy",
				Name:         "InvokeGliteConnectivity",
				Intype: func() macro.Parameters {
					var parameters MeshPreProxyInvokeGliteConnectivityParameters
					return &parameters
				},
				Retype:   func() macro.Returns { var returns MeshPreProxyInvokeGliteConnectivityReturns; return &returns },
				Inbound:  func() macro.Parameters { return new(MeshPreProxyInvokeGliteConnectivityParameters) },
				Outbound: func() macro.Returns { return new(MeshPreProxyInvokeGliteConnectivityReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.glite.connectivity.api"},
				},
			},
			"InvokeHealthcheck": {
				DeclaredKind: (*prsim.PreProxy)(nil),
				TName:        "prsim.PreProxy",
				Name:         "InvokeHealthcheck",
				Intype:       func() macro.Parameters { var parameters MeshPreProxyInvokeHealthcheckParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshPreProxyInvokeHealthcheckReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshPreProxyInvokeHealthcheckParameters) },
				Outbound:     func() macro.Returns { return new(MeshPreProxyInvokeHealthcheckReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.glite.healthcheck.api"},
				},
			},
		},
	})
}

// meshPreProxyMPI is an implementation of PreProxy
type meshPreProxyMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshPreProxyMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshPreProxyMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshPreProxyMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// @MPI("mesh.glite2omega.api")
func (that *meshPreProxyMPI) ProxyInvokeGlite2Omega(ctx context.Context, param interface{}) (interface{}, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["ProxyInvokeGlite2Omega"], param)
	if nil != err {
		x := new(MeshPreProxyProxyInvokeGlite2OmegaReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPreProxyProxyInvokeGlite2OmegaReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPreProxyProxyInvokeGlite2OmegaReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// @MPI("mesh.omega2glite.api")
func (that *meshPreProxyMPI) ProxyInvokeOmega2Glite(ctx context.Context, param interface{}) (interface{}, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["ProxyInvokeOmega2Glite"], param)
	if nil != err {
		x := new(MeshPreProxyProxyInvokeOmega2GliteReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPreProxyProxyInvokeOmega2GliteReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPreProxyProxyInvokeOmega2GliteReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// @MPI("mesh.glite.connectivity.api")
func (that *meshPreProxyMPI) InvokeGliteConnectivity(ctx context.Context, param interface{}) (interface{}, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["InvokeGliteConnectivity"], param)
	if nil != err {
		x := new(MeshPreProxyInvokeGliteConnectivityReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPreProxyInvokeGliteConnectivityReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPreProxyInvokeGliteConnectivityReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// @MPI("mesh.glite.healthcheck.api")
func (that *meshPreProxyMPI) InvokeHealthcheck(ctx context.Context, param interface{}) (interface{}, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["InvokeHealthcheck"], param)
	if nil != err {
		x := new(MeshPreProxyInvokeHealthcheckReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPreProxyInvokeHealthcheckReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPreProxyInvokeHealthcheckReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

type MeshPreProxyProxyInvokeGlite2OmegaParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Param       interface{}       `index:"0" json:"param" xml:"param" yaml:"param"`
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaParameters) GetKind() interface{} {
	return new(MeshPreProxyProxyInvokeGlite2OmegaParameters)
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Param)
	return arguments
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Param = arguments[0].(interface{})
		}
	}
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPreProxyProxyInvokeGlite2OmegaReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content interface{}  `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) GetCode() string {
	return that.Code
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPreProxyProxyInvokeGlite2OmegaReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(interface{})
		}
	}
}

type MeshPreProxyProxyInvokeOmega2GliteParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Param       interface{}       `index:"0" json:"param" xml:"param" yaml:"param"`
}

func (that *MeshPreProxyProxyInvokeOmega2GliteParameters) GetKind() interface{} {
	return new(MeshPreProxyProxyInvokeOmega2GliteParameters)
}

func (that *MeshPreProxyProxyInvokeOmega2GliteParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Param)
	return arguments
}

func (that *MeshPreProxyProxyInvokeOmega2GliteParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Param = arguments[0].(interface{})
		}
	}
}

func (that *MeshPreProxyProxyInvokeOmega2GliteParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPreProxyProxyInvokeOmega2GliteParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPreProxyProxyInvokeOmega2GliteReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content interface{}  `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) GetCode() string {
	return that.Code
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPreProxyProxyInvokeOmega2GliteReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(interface{})
		}
	}
}

type MeshPreProxyInvokeGliteConnectivityParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Param       interface{}       `index:"0" json:"param" xml:"param" yaml:"param"`
}

func (that *MeshPreProxyInvokeGliteConnectivityParameters) GetKind() interface{} {
	return new(MeshPreProxyInvokeGliteConnectivityParameters)
}

func (that *MeshPreProxyInvokeGliteConnectivityParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Param)
	return arguments
}

func (that *MeshPreProxyInvokeGliteConnectivityParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Param = arguments[0].(interface{})
		}
	}
}

func (that *MeshPreProxyInvokeGliteConnectivityParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPreProxyInvokeGliteConnectivityParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPreProxyInvokeGliteConnectivityReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content interface{}  `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) GetCode() string {
	return that.Code
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPreProxyInvokeGliteConnectivityReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(interface{})
		}
	}
}

type MeshPreProxyInvokeHealthcheckParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Param       interface{}       `index:"0" json:"param" xml:"param" yaml:"param"`
}

func (that *MeshPreProxyInvokeHealthcheckParameters) GetKind() interface{} {
	return new(MeshPreProxyInvokeHealthcheckParameters)
}

func (that *MeshPreProxyInvokeHealthcheckParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Param)
	return arguments
}

func (that *MeshPreProxyInvokeHealthcheckParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Param = arguments[0].(interface{})
		}
	}
}

func (that *MeshPreProxyInvokeHealthcheckParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPreProxyInvokeHealthcheckParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPreProxyInvokeHealthcheckReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content interface{}  `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPreProxyInvokeHealthcheckReturns) GetCode() string {
	return that.Code
}

func (that *MeshPreProxyInvokeHealthcheckReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPreProxyInvokeHealthcheckReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPreProxyInvokeHealthcheckReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPreProxyInvokeHealthcheckReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPreProxyInvokeHealthcheckReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPreProxyInvokeHealthcheckReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPreProxyInvokeHealthcheckReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(interface{})
		}
	}
}
