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

	"github.com/be-io/mesh/client/golang/cause"
	"github.com/be-io/mesh/client/golang/macro"
	"github.com/be-io/mesh/client/golang/mpc"
	"github.com/be-io/mesh/client/golang/prsim"
	"github.com/be-io/mesh/client/golang/types"
)

func init() {
	var _ macro.Interface = new(meshTransportMPI)
	macro.Provide((*prsim.Transport)(nil), &meshTransportMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"Open": {
				DeclaredKind: (*prsim.Transport)(nil),
				TName:        "prsim.Transport",
				Name:         "Open",
				Intype:       func() macro.Parameters { var parameters MeshTransportOpenParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshTransportOpenReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshTransportOpenParameters) },
				Outbound:     func() macro.Returns { return new(MeshTransportOpenReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.chan.open"},
				},
			},
			"Close": {
				DeclaredKind: (*prsim.Transport)(nil),
				TName:        "prsim.Transport",
				Name:         "Close",
				Intype:       func() macro.Parameters { var parameters MeshTransportCloseParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshTransportCloseReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshTransportCloseParameters) },
				Outbound:     func() macro.Returns { return new(MeshTransportCloseReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.chan.close"},
				},
			},
			"Roundtrip": {
				DeclaredKind: (*prsim.Transport)(nil),
				TName:        "prsim.Transport",
				Name:         "Roundtrip",
				Intype:       func() macro.Parameters { var parameters MeshTransportRoundtripParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshTransportRoundtripReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshTransportRoundtripParameters) },
				Outbound:     func() macro.Returns { return new(MeshTransportRoundtripReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.chan.roundtrip"},
				},
			},
		},
	})
}

// meshTransportMPI is an implementation of Transport
type meshTransportMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshTransportMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshTransportMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshTransportMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// Open a channel session.
// @MPI("mesh.chan.open")
func (that *meshTransportMPI) Open(ctx context.Context, sessionId string, metadata map[string]string) (prsim.Session, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Open"], sessionId, metadata)
	if nil != err {
		x := new(MeshTransportOpenReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshTransportOpenReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshTransportOpenReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Close the channel.
// @MPI("mesh.chan.close")
func (that *meshTransportMPI) Close(ctx context.Context, timeout types.Duration) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["Close"], timeout)
	return err
}

// Roundtrip with the channel.
// @MPI("mesh.chan.roundtrip")
func (that *meshTransportMPI) Roundtrip(ctx context.Context, payload []byte, metadata map[string]string) ([]byte, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Roundtrip"], payload, metadata)
	if nil != err {
		x := new(MeshTransportRoundtripReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshTransportRoundtripReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshTransportRoundtripReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

type MeshTransportOpenParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	SessionId   string            `index:"0" json:"session_id" xml:"session_id" yaml:"session_id"`
	Metadata    map[string]string `index:"1" json:"metadata" xml:"metadata" yaml:"metadata"`
}

func (that *MeshTransportOpenParameters) GetKind() interface{} {
	return new(MeshTransportOpenParameters)
}

func (that *MeshTransportOpenParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.SessionId)
	arguments = append(arguments, that.Metadata)
	return arguments
}

func (that *MeshTransportOpenParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.SessionId = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Metadata = arguments[1].(map[string]string)
		}
	}
}

func (that *MeshTransportOpenParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshTransportOpenParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshTransportOpenReturns struct {
	Code    string        `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string        `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause  `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content prsim.Session `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshTransportOpenReturns) GetCode() string {
	return that.Code
}

func (that *MeshTransportOpenReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshTransportOpenReturns) GetMessage() string {
	return that.Message
}

func (that *MeshTransportOpenReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshTransportOpenReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshTransportOpenReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshTransportOpenReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshTransportOpenReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(prsim.Session)
		}
	}
}

type MeshTransportCloseParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Timeout     types.Duration    `index:"0" json:"timeout" xml:"timeout" yaml:"timeout"`
}

func (that *MeshTransportCloseParameters) GetKind() interface{} {
	return new(MeshTransportCloseParameters)
}

func (that *MeshTransportCloseParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Timeout)
	return arguments
}

func (that *MeshTransportCloseParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Timeout = arguments[0].(types.Duration)
		}
	}
}

func (that *MeshTransportCloseParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshTransportCloseParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshTransportCloseReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshTransportCloseReturns) GetCode() string {
	return that.Code
}

func (that *MeshTransportCloseReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshTransportCloseReturns) GetMessage() string {
	return that.Message
}

func (that *MeshTransportCloseReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshTransportCloseReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshTransportCloseReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshTransportCloseReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshTransportCloseReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}

type MeshTransportRoundtripParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Payload     []byte            `index:"0" json:"payload" xml:"payload" yaml:"payload"`
	Metadata    map[string]string `index:"1" json:"metadata" xml:"metadata" yaml:"metadata"`
}

func (that *MeshTransportRoundtripParameters) GetKind() interface{} {
	return new(MeshTransportRoundtripParameters)
}

func (that *MeshTransportRoundtripParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Payload)
	arguments = append(arguments, that.Metadata)
	return arguments
}

func (that *MeshTransportRoundtripParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Payload = arguments[0].([]byte)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Metadata = arguments[1].(map[string]string)
		}
	}
}

func (that *MeshTransportRoundtripParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshTransportRoundtripParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshTransportRoundtripReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content []byte       `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshTransportRoundtripReturns) GetCode() string {
	return that.Code
}

func (that *MeshTransportRoundtripReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshTransportRoundtripReturns) GetMessage() string {
	return that.Message
}

func (that *MeshTransportRoundtripReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshTransportRoundtripReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshTransportRoundtripReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshTransportRoundtripReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshTransportRoundtripReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].([]byte)
		}
	}
}
