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
	var _ macro.Interface = new(meshCacheMPI)
	macro.Provide((*prsim.Cache)(nil), &meshCacheMPI{
		invoker: mpc.ServiceProxy.Reference(&macro.Rtt{}),
		methods: map[string]*macro.Method{
			"Get": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Get",
				Intype:       func() macro.Parameters { var parameters MeshCacheGetParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheGetReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheGetParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheGetReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.get"},
				},
			},
			"Put": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Put",
				Intype:       func() macro.Parameters { var parameters MeshCachePutParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCachePutReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCachePutParameters) },
				Outbound:     func() macro.Returns { return new(MeshCachePutReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.put"},
				},
			},
			"Remove": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Remove",
				Intype:       func() macro.Parameters { var parameters MeshCacheRemoveParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheRemoveReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheRemoveParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheRemoveReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.remove"},
				},
			},
			"Incr": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Incr",
				Intype:       func() macro.Parameters { var parameters MeshCacheIncrParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheIncrReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheIncrParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheIncrReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.incr"},
				},
			},
			"Decr": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Decr",
				Intype:       func() macro.Parameters { var parameters MeshCacheDecrParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheDecrReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheDecrParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheDecrReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.decr"},
				},
			},
			"Keys": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "Keys",
				Intype:       func() macro.Parameters { var parameters MeshCacheKeysParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheKeysReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheKeysParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheKeysReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.keys"},
				},
			},
			"HGet": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "HGet",
				Intype:       func() macro.Parameters { var parameters MeshCacheHGetParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheHGetReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheHGetParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheHGetReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.hget"},
				},
			},
			"HSet": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "HSet",
				Intype:       func() macro.Parameters { var parameters MeshCacheHSetParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheHSetReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheHSetParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheHSetReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.hset"},
				},
			},
			"HDel": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "HDel",
				Intype:       func() macro.Parameters { var parameters MeshCacheHDelParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheHDelReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheHDelParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheHDelReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.hdel"},
				},
			},
			"HKeys": {
				DeclaredKind: (*prsim.Cache)(nil),
				TName:        "prsim.Cache",
				Name:         "HKeys",
				Intype:       func() macro.Parameters { var parameters MeshCacheHKeysParameters; return &parameters },
				Retype:       func() macro.Returns { var returns MeshCacheHKeysReturns; return &returns },
				Inbound:      func() macro.Parameters { return new(MeshCacheHKeysParameters) },
				Outbound:     func() macro.Returns { return new(MeshCacheHKeysReturns) },
				MPI: &macro.MPIAnnotation{
					Meta: &macro.Rtt{Name: "mesh.cache.hkeys"},
				},
			},
		},
	})
}

// meshCacheMPI is an implementation of Cache
type meshCacheMPI struct {
	invoker macro.Caller
	methods map[string]*macro.Method
}

func (that *meshCacheMPI) Att() *macro.Att {
	return &macro.Att{Name: macro.MeshMPI}
}

func (that *meshCacheMPI) Rtt() *macro.Rtt {
	return &macro.Rtt{Name: macro.MeshMPI}
}

func (that *meshCacheMPI) GetMethods() map[string]*macro.Method {
	return that.methods
}

// Get the value from cache.
// @MPI("mesh.cache.get")
func (that *meshCacheMPI) Get(ctx context.Context, key string) (*types.CacheEntity, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Get"], key)
	if nil != err {
		x := new(MeshCacheGetReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheGetReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheGetReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Put the value to cache.
// @MPI("mesh.cache.put")
func (that *meshCacheMPI) Put(ctx context.Context, cell *types.CacheEntity) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["Put"], cell)
	return err
}

// Remove the cache value.
// @MPI("mesh.cache.remove")
func (that *meshCacheMPI) Remove(ctx context.Context, key string) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["Remove"], key)
	return err
}

// Incr the cache of expire time.
// @MPI("mesh.cache.incr")
func (that *meshCacheMPI) Incr(ctx context.Context, key string, value int64) (int64, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Incr"], key, value)
	if nil != err {
		x := new(MeshCacheIncrReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheIncrReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheIncrReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Decr the cache of expire time.
// @MPI("mesh.cache.decr")
func (that *meshCacheMPI) Decr(ctx context.Context, key string, value int64) (int64, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Decr"], key, value)
	if nil != err {
		x := new(MeshCacheDecrReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheDecrReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheDecrReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// Keys the cache key set.
// @MPI("mesh.cache.keys")
func (that *meshCacheMPI) Keys(ctx context.Context, pattern string) ([]string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["Keys"], pattern)
	if nil != err {
		x := new(MeshCacheKeysReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheKeysReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheKeysReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// HGet get value in hash
// @MPI("mesh.cache.hget")
func (that *meshCacheMPI) HGet(ctx context.Context, key string, name string) (*types.CacheEntity, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["HGet"], key, name)
	if nil != err {
		x := new(MeshCacheHGetReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheHGetReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheHGetReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

// HSet put value in hash
// @MPI("mesh.cache.hset")
func (that *meshCacheMPI) HSet(ctx context.Context, key string, cell *types.CacheEntity) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["HSet"], key, cell)
	return err
}

// HDel put value in hash
// @MPI("mesh.cache.hdel")
func (that *meshCacheMPI) HDel(ctx context.Context, key string, name string) error {
	_, err := that.invoker.Call(ctx, that.invoker, that.methods["HDel"], key, name)
	return err
}

// HKeys get the hash keys
// @MPI("mesh.cache.hkeys")
func (that *meshCacheMPI) HKeys(ctx context.Context, key string) ([]string, error) {
	ret, err := that.invoker.Call(ctx, that.invoker, that.methods["HKeys"], key)
	if nil != err {
		x := new(MeshCacheHKeysReturns)
		return x.Content, err
	}
	x, ok := ret.(*MeshCacheHKeysReturns)
	if ok {
		return x.Content, err
	}
	x = new(MeshCacheHKeysReturns)
	return x.Content, cause.Errorf("Cant resolve response ")
}

type MeshCacheGetParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
}

func (that *MeshCacheGetParameters) GetKind() interface{} {
	return new(MeshCacheGetParameters)
}

func (that *MeshCacheGetParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	return arguments
}

func (that *MeshCacheGetParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
	}
}

func (that *MeshCacheGetParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheGetParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheGetReturns struct {
	Code    string             `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string             `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause       `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content *types.CacheEntity `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheGetReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheGetReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheGetReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheGetReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheGetReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheGetReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheGetReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheGetReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(*types.CacheEntity)
		}
	}
}

type MeshCachePutParameters struct {
	Attachments map[string]string  `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Cell        *types.CacheEntity `index:"0" json:"cell" xml:"cell" yaml:"cell"`
}

func (that *MeshCachePutParameters) GetKind() interface{} {
	return new(MeshCachePutParameters)
}

func (that *MeshCachePutParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Cell)
	return arguments
}

func (that *MeshCachePutParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Cell = arguments[0].(*types.CacheEntity)
		}
	}
}

func (that *MeshCachePutParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCachePutParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCachePutReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshCachePutReturns) GetCode() string {
	return that.Code
}

func (that *MeshCachePutReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCachePutReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCachePutReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCachePutReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCachePutReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCachePutReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshCachePutReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}

type MeshCacheRemoveParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
}

func (that *MeshCacheRemoveParameters) GetKind() interface{} {
	return new(MeshCacheRemoveParameters)
}

func (that *MeshCacheRemoveParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	return arguments
}

func (that *MeshCacheRemoveParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
	}
}

func (that *MeshCacheRemoveParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheRemoveParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheRemoveReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshCacheRemoveReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheRemoveReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheRemoveReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheRemoveReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheRemoveReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheRemoveReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheRemoveReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshCacheRemoveReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}

type MeshCacheIncrParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
	Value       int64             `index:"1" json:"value" xml:"value" yaml:"value"`
}

func (that *MeshCacheIncrParameters) GetKind() interface{} {
	return new(MeshCacheIncrParameters)
}

func (that *MeshCacheIncrParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	arguments = append(arguments, that.Value)
	return arguments
}

func (that *MeshCacheIncrParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Value = arguments[1].(int64)
		}
	}
}

func (that *MeshCacheIncrParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheIncrParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheIncrReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content int64        `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheIncrReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheIncrReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheIncrReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheIncrReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheIncrReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheIncrReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheIncrReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheIncrReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(int64)
		}
	}
}

type MeshCacheDecrParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
	Value       int64             `index:"1" json:"value" xml:"value" yaml:"value"`
}

func (that *MeshCacheDecrParameters) GetKind() interface{} {
	return new(MeshCacheDecrParameters)
}

func (that *MeshCacheDecrParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	arguments = append(arguments, that.Value)
	return arguments
}

func (that *MeshCacheDecrParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Value = arguments[1].(int64)
		}
	}
}

func (that *MeshCacheDecrParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheDecrParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheDecrReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content int64        `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheDecrReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheDecrReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheDecrReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheDecrReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheDecrReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheDecrReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheDecrReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheDecrReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(int64)
		}
	}
}

type MeshCacheKeysParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Pattern     string            `index:"0" json:"pattern" xml:"pattern" yaml:"pattern"`
}

func (that *MeshCacheKeysParameters) GetKind() interface{} {
	return new(MeshCacheKeysParameters)
}

func (that *MeshCacheKeysParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Pattern)
	return arguments
}

func (that *MeshCacheKeysParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Pattern = arguments[0].(string)
		}
	}
}

func (that *MeshCacheKeysParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheKeysParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheKeysReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content []string     `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheKeysReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheKeysReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheKeysReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheKeysReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheKeysReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheKeysReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheKeysReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheKeysReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].([]string)
		}
	}
}

type MeshCacheHGetParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
	Name        string            `index:"1" json:"name" xml:"name" yaml:"name"`
}

func (that *MeshCacheHGetParameters) GetKind() interface{} {
	return new(MeshCacheHGetParameters)
}

func (that *MeshCacheHGetParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	arguments = append(arguments, that.Name)
	return arguments
}

func (that *MeshCacheHGetParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Name = arguments[1].(string)
		}
	}
}

func (that *MeshCacheHGetParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheHGetParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheHGetReturns struct {
	Code    string             `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string             `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause       `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content *types.CacheEntity `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheHGetReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheHGetReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheHGetReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheHGetReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheHGetReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheHGetReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheHGetReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheHGetReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].(*types.CacheEntity)
		}
	}
}

type MeshCacheHSetParameters struct {
	Attachments map[string]string  `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string             `index:"0" json:"key" xml:"key" yaml:"key"`
	Cell        *types.CacheEntity `index:"1" json:"cell" xml:"cell" yaml:"cell"`
}

func (that *MeshCacheHSetParameters) GetKind() interface{} {
	return new(MeshCacheHSetParameters)
}

func (that *MeshCacheHSetParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	arguments = append(arguments, that.Cell)
	return arguments
}

func (that *MeshCacheHSetParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Cell = arguments[1].(*types.CacheEntity)
		}
	}
}

func (that *MeshCacheHSetParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheHSetParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheHSetReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshCacheHSetReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheHSetReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheHSetReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheHSetReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheHSetReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheHSetReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheHSetReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshCacheHSetReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}

type MeshCacheHDelParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
	Name        string            `index:"1" json:"name" xml:"name" yaml:"name"`
}

func (that *MeshCacheHDelParameters) GetKind() interface{} {
	return new(MeshCacheHDelParameters)
}

func (that *MeshCacheHDelParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	arguments = append(arguments, that.Name)
	return arguments
}

func (that *MeshCacheHDelParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
		if len(arguments) > 1 && nil != arguments[1] {
			that.Name = arguments[1].(string)
		}
	}
}

func (that *MeshCacheHDelParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheHDelParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheHDelReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
}

func (that *MeshCacheHDelReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheHDelReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheHDelReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheHDelReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheHDelReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheHDelReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheHDelReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	return arguments
}

func (that *MeshCacheHDelReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
	}
}

type MeshCacheHKeysParameters struct {
	Attachments map[string]string `index:"-1" json:"attachments" xml:"attachments" yaml:"attachments"`
	Key         string            `index:"0" json:"key" xml:"key" yaml:"key"`
}

func (that *MeshCacheHKeysParameters) GetKind() interface{} {
	return new(MeshCacheHKeysParameters)
}

func (that *MeshCacheHKeysParameters) GetArguments(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Key)
	return arguments
}

func (that *MeshCacheHKeysParameters) SetArguments(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Key = arguments[0].(string)
		}
	}
}

func (that *MeshCacheHKeysParameters) GetAttachments(ctx context.Context) map[string]string {
	return that.Attachments
}

func (that *MeshCacheHKeysParameters) SetAttachments(ctx context.Context, attachments map[string]string) {
	that.Attachments = attachments
}

type MeshCacheHKeysReturns struct {
	Code    string       `index:"0" json:"code" xml:"code" yaml:"code" comment:"Result code"`
	Message string       `index:"5" json:"message" xml:"message" yaml:"message" comment:"Result message"`
	Cause   *macro.Cause `index:"10" json:"cause" xml:"cause" yaml:"cause" comment:"Service cause stacktrace"`
	Content []string     `index:"15" json:"content" xml:"content" yaml:"content"`
}

func (that *MeshCacheHKeysReturns) GetCode() string {
	return that.Code
}

func (that *MeshCacheHKeysReturns) SetCode(code string) {
	that.Code = code
}

func (that *MeshCacheHKeysReturns) GetMessage() string {
	return that.Message
}

func (that *MeshCacheHKeysReturns) SetMessage(message string) {
	that.Message = message
}

func (that *MeshCacheHKeysReturns) GetCause(ctx context.Context) *macro.Cause {
	return that.Cause
}

func (that *MeshCacheHKeysReturns) SetCause(ctx context.Context, cause *macro.Cause) {
	that.Cause = cause
}

func (that *MeshCacheHKeysReturns) GetContent(ctx context.Context) []interface{} {
	var arguments []interface{}
	arguments = append(arguments, that.Content)
	return arguments
}

func (that *MeshCacheHKeysReturns) SetContent(ctx context.Context, arguments ...interface{}) {
	if len(arguments) > 0 {
		if len(arguments) > 0 && nil != arguments[0] {
			that.Content = arguments[0].([]string)
		}
	}
}
