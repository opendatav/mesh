/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package metabase

import (
	"bytes"
	"context"
	"git.firmer.tech/firmer/mesh/client/golang/mesh"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/client/golang/tool"
	"github.com/opendatav/mesh/client/golang/types"
	"github.com/opendatav/mesh/plugin/metabase/orm"
	"time"
)

func init() {
	var _ prsim.KV = omegaKv
	macro.Provide(prsim.IKV, omegaKv)
}

var omegaKv = new(kv)

type kv struct {
}

func (that *kv) Att() *macro.Att {
	return &macro.Att{Name: Name}
}

func (that *kv) Get(ctx context.Context, key string) (*types.Entity, error) {
	return WithoutTxReturn(ctx, func(ctx context.Context, session orm.IO) (*types.Entity, error) {
		set, err := session.Kv().SelectMeshKvByKey(ctx, key)
		if IsNoRow(err) {
			return new(types.Entity).AsEmpty(), nil
		}
		if nil != err {
			return nil, cause.Error(err)
		}
		var entity types.Entity
		if _, err = aware.Codec.Decode(bytes.NewBufferString(set.Value), &entity); nil != err {
			return nil, cause.Error(err)
		}
		return &entity, nil
	})
}

func (that *kv) Put(ctx context.Context, key string, value *types.Entity) error {
	return WithTx(ctx, func(ctx context.Context, session orm.IO) error {
		return that.PutWithSession(ctx, session, key, value)
	})
}

func (that *kv) PutWithSession(ctx context.Context, session orm.IO, key string, value *types.Entity) error {
	entity, err := aware.Codec.EncodeString(value)
	if nil != err {
		return cause.Error(err)
	}
	mtx := mesh.ContextWith(ctx)
	operator := tool.Anyone(mtx.GetAttachments()["omega.user.username"], mtx.GetAttachments()["omega.inst.name"])
	exist, err := session.Kv().SelectMeshKvByKey(ctx, key)
	if IsNoRow(err) || nil == exist {
		_, err = session.Kv().InsertMeshKv(ctx, &orm.MeshKv{
			Key:      key,
			Value:    entity,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
			CreateBy: operator,
			UpdateBy: operator,
		})
		return cause.Error(err)
	}
	if IsNoRow(err) {
		return cause.Error(err)
	}
	_, err = session.Kv().UpdateMeshKvByKey(ctx, &orm.MeshKv{
		Key:      key,
		Value:    entity,
		CreateAt: exist.CreateAt,
		UpdateAt: time.Now(),
		CreateBy: exist.CreateBy,
		UpdateBy: operator,
	})
	return cause.Error(err)
}

func (that *kv) Remove(ctx context.Context, key string) error {
	return WithTx(ctx, func(ctx context.Context, session orm.IO) error {
		_, err := session.Kv().DeleteMeshKvByKey(ctx, key)
		return cause.Error(err)
	})
}

func (that *kv) Keys(ctx context.Context, pattern string) ([]string, error) {
	return WithoutTxReturn(ctx, func(ctx context.Context, session orm.IO) ([]string, error) {
		keys, err := session.Kv().SelectKeys(ctx, pattern)
		if IsNoRow(err) {
			return nil, nil
		}
		return keys, nil
	})
}

func (that *kv) Index(ctx context.Context, index *types.Paging) (*types.Page[*types.Entry], error) {
	return WithoutTxReturn(ctx, func(ctx context.Context, session orm.IO) (*types.Page[*types.Entry], error) {
		rets, err := session.Kv().Index(ctx, index.Index, index.Limit)
		if nil != err {
			return nil, cause.Error(err)
		}
		entries := make([]*types.Entry, len(rets.Data))
		for idx, ret := range rets.Data {
			var entity types.Entity
			if _, err = aware.Codec.Decode(bytes.NewBufferString(ret.Value), &entity); nil != err {
				return nil, cause.Error(err)
			}
			entries[idx] = &types.Entry{
				Key:      ret.Key,
				Value:    &entity,
				UpdateAt: types.Time(ret.UpdateAt),
			}
		}
		return types.Reset(index, rets.Total, entries), nil
	})
}
