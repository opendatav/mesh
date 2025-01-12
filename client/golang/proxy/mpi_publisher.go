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
	var _ macro.Interface = new(meshPublisherMPI)
	macro.Provide((*prsim.Publisher)(nil), &meshPublisherMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"Publish": {
				DeclaredKind: (*prsim.Publisher)(nil),
				TName:        "prsim.Publisher",
				Name:         "Publish",
				Intype:       func() macro.Parameters { var parameters MeshPublisherPublishParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshPublisherPublishReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshPublisherPublishParameters) },
				Outbound:     func() macro.Returns { return new(MeshPublisherPublishReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.queue.publish"},
				},
			},
			"Broadcast": {
				DeclaredKind: (*prsim.Publisher)(nil),
				TName:        "prsim.Publisher",
				Name:         "Broadcast",
				Intype:       func() macro.Parameters { var parameters MeshPublisherBroadcastParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshPublisherBroadcastReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshPublisherBroadcastParameters) },
				Outbound:     func() macro.Returns { return new(MeshPublisherBroadcastReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.queue.multicast"},
				},
			},
		},
	})
}

// meshPublisherMPI is an implementation of Publisher
type meshPublisherMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshPublisherMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshPublisherMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshPublisherMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// Publish
// @MPI("mesh.queue.publish")
func (that *meshPublisherMPI) Publish(ctx context.Context, events []*types.Event) ([]string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Publish"], events)
	if nil != err {
		x := new(MeshPublisherPublishReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPublisherPublishReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPublisherPublishReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Broadcast
// @MPI("mesh.queue.multicast")
func (that *meshPublisherMPI) Broadcast(ctx context.Context, events []*types.Event) ([]string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Broadcast"], events)
	if nil != err {
		x := new(MeshPublisherBroadcastReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshPublisherBroadcastReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshPublisherBroadcastReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

type MeshPublisherPublishParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Events      []*types.Event    `index:"0" json:"events" xml:"events" yaml:"events"`
}

func (that *MeshPublisherPublishParameters) GetKind() interface{} {
	return new(MeshPublisherPublishParameters)
}

func (that *MeshPublisherPublishParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Events)
	return arguments
}

func (that *MeshPublisherPublishParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Events = arguments[0].([]*types.Event)
		}
	}
}

func (that *MeshPublisherPublishParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPublisherPublishParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPublisherPublishReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content []string     `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPublisherPublishReturns) GetCode() string {
	return that.Code
}

func (that *MeshPublisherPublishReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPublisherPublishReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPublisherPublishReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPublisherPublishReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPublisherPublishReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPublisherPublishReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPublisherPublishReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].([]string)
		}
	}
}

type MeshPublisherBroadcastParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Events      []*types.Event    `index:"0" json:"events" xml:"events" yaml:"events"`
}

func (that *MeshPublisherBroadcastParameters) GetKind() interface{} {
	return new(MeshPublisherBroadcastParameters)
}

func (that *MeshPublisherBroadcastParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Events)
	return arguments
}

func (that *MeshPublisherBroadcastParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Events = arguments[0].([]*types.Event)
		}
	}
}

func (that *MeshPublisherBroadcastParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshPublisherBroadcastParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshPublisherBroadcastReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content []string     `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshPublisherBroadcastReturns) GetCode() string {
	return that.Code
}

func (that *MeshPublisherBroadcastReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshPublisherBroadcastReturns) GetMessage() string {
	return that.Message
}

func (that *MeshPublisherBroadcastReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshPublisherBroadcastReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshPublisherBroadcastReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshPublisherBroadcastReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshPublisherBroadcastReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].([]string)
		}
	}
}
