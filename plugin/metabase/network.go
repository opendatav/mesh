/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package metabase

import (
	"context"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/log"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/system"
	"github.com/opendatav/mesh/client/golang/tool"
	"github.com/opendatav/mesh/client/golang/types"
	"github.com/opendatav/mesh/plugin/metabase/orm"
	"time"
)

func init() {
	var _ prsim.Network = lan
	macro.Provide(prsim.INetwork, lan)
}

var lan = new(network)

// 原来的组网管理改成节点管理。原来的操作按钮组网改成手动组网。
// 里面新增一个状态位：是否已组网，和原有的是否联通区分开
// 新增一个标记列：节点来源（手动组网、自动发现、入网申请）
// 操作里面对组网状态进行判断，未组网的，新增一个组网申请操作，入网申请状态，新增一个申请审批
//
// 变更点：
// 1、组网管理升级成节点管理，组网只是节点管理中的一个操作。
// 2、新增组网申请与审批两个操作。
// 3、支持手动组网和自动发现组网，节点列表里面三种来源：手动组网、自动发现、入网申请
type network struct {
}

func (that *network) Att() *macro.Att {
	return &macro.Att{Name: Name}
}

func (that *network) GetEnviron(ctx context.Context) (*types.Environ, error) {
	return system.Environ.Get(), nil
}

func (that *network) Accessible(ctx context.Context, route *types.Route) (bool, error) {
	return true, nil
}

func (that *network) Refresh(ctx context.Context, routes []*types.Route) error {
	return WithoutTx(ctx, func(ctx context.Context, session orm.IO) error {
		for _, route := range routes {
			if "" == route.NodeId {
				continue
			}
			if (route.Status & int32(types.Removed)) == int32(types.Removed) {
				if _, err := session.Edge().DeleteEdge(ctx, route.NodeId); nil != err {
					return cause.Error(err)
				}
				continue
			}
			extra, err := new(types.RouteAttribute).Encode(ctx, route, aware.Codec)
			if nil != err {
				log.Warn(ctx, "Encode:%s", err.Error())
				return cause.Error(err)
			}
			prev, err := session.Edge().SelectEdgeByNodeId(ctx, route.NodeId)
			if nil != err && !IsNoRow(err) {
				return cause.Error(err)
			}
			certification := new(types.RouteCertificate)
			certification.Decode(ctx, prev.Certificate, aware.Codec)
			edge := &orm.Edge{
				NodeId:      route.NodeId,
				InstId:      tool.Anyone(route.InstId, prev.InstId),
				InstName:    tool.Anyone(route.Name, route.InstName, prev.InstName),
				Address:     tool.Anyone(route.URC().String(), prev.Address),
				Describe:    tool.Anyone(route.Describe, prev.Describe),
				Certificate: certification.Override0(route).Encode(ctx, aware.Codec),
				Status:      int64(types.Weaved.OrElse(int32(prev.Status), types.Connected.OrElse(int32(prev.Status), route.Status))),
				Version:     prev.Version + 1,
				AuthCode:    tool.Anyone(route.AuthCode, prev.AuthCode),
				Extra:       tool.Anyone(extra, prev.Extra),
				ExpireAt:    time.UnixMilli(tool.Ternary(route.ExpireAt < 1, time.Now().AddDate(10, 0, 0).UnixMilli(), route.ExpireAt)),
				CreateAt:    tool.Anyone(prev.CreateAt, time.Now()),
				UpdateAt:    time.Now(),
				CreateBy:    tool.Anyone(prev.CreateBy, route.CreateBy),
				UpdateBy:    tool.Anyone(route.UpdateBy, prev.UpdateBy),
				Group:       route.Group,
				StaticIp:    tool.Anyone(route.StaticIP, prev.StaticIp),
				PublicIp:    tool.Anyone(route.PublicIP, prev.PublicIp),
				Requests:    int64(tool.Anyone(route.Requests, int(prev.Requests))),
			}
			if IsNoRow(err) {
				if _, err = session.Edge().InsertEdge(ctx, edge); nil != err {
					return cause.Error(err)
				}
			}
			if _, err = session.Edge().UpdateEdgeByNodeId(ctx, edge); nil != err {
				return cause.Error(err)
			}
		}
		return nil
	})
}

func (that *network) GetRoute(ctx context.Context, nodeId string) (*types.Route, error) {
	return WithoutTxReturn(ctx, func(ctx context.Context, session orm.IO) (*types.Route, error) {
		edge, err := session.Edge().SelectByIds(ctx, nodeId, nodeId)
		if nil == err {
			return that.ToRoute(ctx, edge), nil
		}
		if !IsNoRow(err) {
			return nil, cause.Error(err)
		}
		return nil, nil
	})
}

func (that *network) GetRoutes(ctx context.Context) ([]*types.Route, error) {
	return WithoutTxReturn(ctx, func(ctx context.Context, session orm.IO) ([]*types.Route, error) {
		edges, err := session.Edge().SelectAll(ctx)
		if nil != err {
			return nil, cause.Error(err)
		}
		return that.ToRoutes(ctx, edges), nil
	})
}

func (that *network) GetDomains(ctx context.Context, kind string) ([]*types.Domain, error) {
	return nil, nil
}

func (that *network) PutDomains(ctx context.Context, kind string, domains []*types.Domain) error {
	return nil
}

func (that *network) Weave(ctx context.Context, route *types.Route) error {
	route.Status = route.Status | int32(types.FromWeaving)
	return cause.Error(that.Refresh(ctx, []*types.Route{route}))
}

func (that *network) Ack(ctx context.Context, route *types.Route) error {
	route.Status = route.Status | int32(types.Weaved)
	return cause.Error(that.Refresh(ctx, []*types.Route{route}))
}

func (that *network) Disable(ctx context.Context, nodeId string) error {
	edge, err := that.GetRoute(ctx, nodeId)
	if nil != err {
		return cause.Error(err)
	}
	if nil == edge {
		return cause.Error(cause.Errorable(cause.NetNotWeave))
	}
	edge.Status = edge.Status | int32(types.Disabled)
	return cause.Error(that.Refresh(ctx, []*types.Route{edge}))
}

func (that *network) Enable(ctx context.Context, nodeId string) error {
	edge, err := that.GetRoute(ctx, nodeId)
	if nil != err {
		return cause.Error(err)
	}
	if nil == edge {
		return cause.Error(cause.Errorable(cause.NetNotWeave))
	}
	edge.Status = edge.Status & int32(^types.Disabled)
	return cause.Error(that.Refresh(ctx, []*types.Route{edge}))
}

func (that *network) Index(ctx context.Context, index *types.Paging) (*types.Page[*types.Route], error) {
	return WithTxReturn(ctx, func(ctx context.Context, session orm.IO) (*types.Page[*types.Route], error) {
		rs, err := session.Edge().IndexEdge(ctx, index.Index, index.Limit)
		if nil != err {
			return nil, cause.Error(err)
		}
		return types.Reset(index, int64(rs.Total), that.ToRoutes(ctx, rs.Data)), nil
	})
}

func (that *network) Version(ctx context.Context, nodeId string) (*types.Versions, error) {
	return &types.Versions{Version: prsim.Version}, nil
}

func (that *network) Instx(ctx context.Context, index *types.Paging) (*types.Page[*types.Institution], error) {
	return nil, nil
}

func (that *network) Instr(ctx context.Context, institutions []*types.Institution) error {
	return nil
}

func (that *network) Ally(ctx context.Context, nodeIds []string) error {
	return nil
}

func (that *network) Disband(ctx context.Context, nodeIds []string) error {
	return nil
}

func (that *network) Assert(ctx context.Context, feature string, nodeIds []string) (bool, error) {
	return false, nil
}

func (that *network) ToRoutes(ctx context.Context, edges []*orm.Edge) []*types.Route {
	var es []*types.Route
	for _, edge := range edges {
		es = append(es, that.ToRoute(ctx, edge))
	}
	return es
}

func (that *network) ToRoute(ctx context.Context, route *orm.Edge) *types.Route {
	certification := new(types.RouteCertificate)
	certification.Decode(ctx, route.Certificate, aware.Codec)
	attribute := new(types.RouteAttribute).Decode(ctx, route.Extra, aware.Codec)
	return &types.Route{
		NodeId:      route.NodeId,
		InstId:      route.InstId,
		Name:        route.InstName,
		InstName:    route.InstName,
		Address:     route.Address,
		Describe:    route.Describe,
		HostRoot:    certification.HostRoot,
		HostKey:     certification.HostKey,
		HostCrt:     certification.HostCrt,
		GuestRoot:   certification.GuestRoot,
		GuestKey:    certification.GuestKey,
		GuestCrt:    certification.GuestCrt,
		Status:      int32(route.Status),
		Version:     int32(route.Version),
		AuthCode:    route.AuthCode,
		ExpireAt:    route.ExpireAt.UnixMilli(),
		Extra:       attribute.Compat(ctx, aware.Codec),
		CreateAt:    route.CreateAt.Format(log.DateFormat),
		UpdateAt:    route.UpdateAt.Format(log.DateFormat),
		CreateBy:    route.CreateBy,
		UpdateBy:    route.UpdateBy,
		Group:       route.Group,
		Upstream:    attribute.Upstream,
		Downstream:  attribute.Downstream,
		StaticIP:    tool.Anyone(route.StaticIp, attribute.StaticIP),
		PublicIP:    route.PublicIp,
		Requests:    int(route.Requests),
		Proxy:       attribute.Proxy,
		Concurrency: attribute.Concurrency,
	}
}
