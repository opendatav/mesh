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

	"github.com/be-io/mesh/client/golang/macro"
	"github.com/be-io/mesh/client/golang/mpc"
	"github.com/be-io/mesh/client/golang/prsim"
	"github.com/be-io/mesh/client/golang/types"
)

func init() {
	var _ macro.Interface = new(meshSubscriberMPI)
	macro.Provide((*prsim.Subscriber)(nil), &meshSubscriberMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"Subscribe": {
				DeclaredKind: (*prsim.Subscriber)(nil),
				TName:        "prsim.Subscriber",
				Name:         "Subscribe",
				Intype:       func() macro.Parameters { var parameters MeshSubscriberSubscribeParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshSubscriberSubscribeReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshSubscriberSubscribeParameters) },
				Outbound:     func() macro.Returns { return new(MeshSubscriberSubscribeReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.queue.subscribe"},
				},
			},
		},
	})
}

// meshSubscriberMPI is an implementation of Subscriber
type meshSubscriberMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshSubscriberMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshSubscriberMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshSubscriberMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// Subscribe the event with {@link com.be.mesh.client.annotate.Bindings} or {@link com.be.mesh.client.annotate.Binding}
// @MPI("mesh.queue.subscribe")
func (that *meshSubscriberMPI) Subscribe(ctx context.Context, event *types.Event) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["Subscribe"], event)
	return err
}

type MeshSubscriberSubscribeParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Event       *types.Event      `index:"0" json:"event" xml:"event" yaml:"event"`
}

func (that *MeshSubscriberSubscribeParameters) GetKind() interface{} {
	return new(MeshSubscriberSubscribeParameters)
}

func (that *MeshSubscriberSubscribeParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Event)
	return arguments
}

func (that *MeshSubscriberSubscribeParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Event = arguments[0].(*types.Event)
		}
	}
}

func (that *MeshSubscriberSubscribeParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshSubscriberSubscribeParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshSubscriberSubscribeReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshSubscriberSubscribeReturns) GetCode() string {
	return that.Code
}

func (that *MeshSubscriberSubscribeReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshSubscriberSubscribeReturns) GetMessage() string {
	return that.Message
}

func (that *MeshSubscriberSubscribeReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshSubscriberSubscribeReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshSubscriberSubscribeReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshSubscriberSubscribeReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshSubscriberSubscribeReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}
