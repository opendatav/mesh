/*
 * Copyright (c) 2000, 2023, trustbe and/or its affiliates. All rights reserved.
 * TRUSTBE PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package metabase

import (
	"context"
	"git.firmer.tech/firmer/mesh/client/golang/mesh"
	"git.firmer.tech/firmer/mesh/client/golang/psi"
	"github.com/opendatav/mesh/client/golang/cause"
	"github.com/opendatav/mesh/client/golang/macro"
	"github.com/opendatav/mesh/client/golang/prsim"
	"github.com/opendatav/mesh/plugin/metabase/orm"
	"time"
)

func init() {
	var ss = &sequence{
		s: &prsim.SyncSequence[orm.IO]{
			Syncer:   new(sequenceSyncer),
			Macro:    &macro.Att{Name: Name},
			Sections: map[string]chan string{}},
	}
	var _ psi.Sequence = ss
	macro.Provide(prsim.ISequence, ss)
}

type sequence struct {
	s psi.Sequence
}

func (that *sequence) Att() *macro.Att {
	return &macro.Att{Name: Name}
}

func (that *sequence) Next(ctx context.Context, kind string, length int) (string, error) {
	if len(kind) < 2 {
		return "", cause.Errorf("Unexpected kind %s, at least two chars", kind)
	}
	idc, err := mesh.Zone(ctx)
	if nil != err {
		return "", cause.Error(err)
	}
	seq, err := that.s.Next(ctx, kind, 8)
	if nil != err {
		return "", cause.Error(err)
	}
	id := &mesh.ID{
		Prefix:  kind[0:2],
		Reserve: "00",
		SEQ:     seq,
		MDC:     idc,
	}
	return id.String(), nil
}

func (that *sequence) Section(ctx context.Context, kind string, size int, length int) ([]string, error) {
	return that.s.Section(ctx, kind, size, length)
}

type sequenceSyncer struct {
}

func (that *sequenceSyncer) Tx(ctx context.Context, tx func(session orm.IO) ([]string, error)) ([]string, error) {
	return WithTxReturn[[]string](ctx, func(ctx context.Context, session orm.IO) ([]string, error) {
		return tx(session)
	})
}

func (that *sequenceSyncer) Incr(ctx context.Context, session orm.IO, kind string) (*prsim.Section, error) {
	seq, err := session.Sequence().SelectMeshSeqByKind(ctx, kind)
	if nil != err {
		return nil, cause.Error(err)
	}
	return &prsim.Section{
		Kind:    seq.Kind,
		Min:     seq.Min,
		Max:     seq.Max,
		Size:    int32(seq.Size),
		Length:  int32(seq.Length),
		Version: int32(seq.Version),
	}, nil
}

func (that *sequenceSyncer) Init(ctx context.Context, session orm.IO, section *prsim.Section) error {
	_, err := session.Sequence().InsertMeshSeq(ctx, &orm.MeshSeq{
		Kind:     section.Kind,
		Min:      section.Min,
		Max:      section.Max,
		Size:     int64(section.Size),
		Length:   int64(section.Length),
		Status:   0,
		Version:  int64(section.Version),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	return cause.Error(err)
}

func (that *sequenceSyncer) Sync(ctx context.Context, session orm.IO, section *prsim.Section) error {
	rows, err := session.Sequence().SetMin(ctx, section.Min+int64(section.Size), section.Kind, int64(section.Version))
	if nil != err {
		return cause.Error(err)
	}
	if rows < 1 {
		return cause.Errorf("Lock sequence fail.")
	}
	return nil
}
