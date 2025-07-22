/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package metabase

import (
	"context"
	"errors"
	"git.firmer.tech/firmer/mesh/client/golang/cause"
	"git.firmer.tech/firmer/mesh/client/golang/orm/dao"
	"github.com/go-sql-driver/mysql"
	postgresql "github.com/jackc/pgx/v5/pgconn"
	"github.com/opendatav/mesh/plugin/metabase/orm"
)

func WithTxReturn[T any](ctx context.Context, closure func(ctx context.Context, session orm.IO) (T, error)) (T, error) {
	return dao.WithTx[orm.IO](ctx, func(ctx context.Context, db orm.IO) (T, error) {
		return closure(ctx, db)
	})
}

func WithTx(ctx context.Context, closure func(ctx context.Context, session orm.IO) error) (err error) {
	return dao.WithTx0[orm.IO](ctx, func(ctx context.Context, db orm.IO) error {
		return closure(ctx, db)
	})
}

func WithoutTxReturn[T any](ctx context.Context, closure func(ctx context.Context, session orm.IO) (T, error)) (T, error) {
	return dao.WithoutTx[orm.IO](ctx, func(ctx context.Context, db orm.IO) (T, error) {
		return closure(ctx, db)
	})
}

func WithoutTx(ctx context.Context, closure func(ctx context.Context, session orm.IO) error) (err error) {
	return dao.WithoutTx0[orm.IO](ctx, func(ctx context.Context, db orm.IO) error {
		return closure(ctx, db)
	})
}

func IsDuplicateKey(err error) bool {
	if nil == err {
		return false
	}
	var errMySQL *mysql.MySQLError
	if errors.As(cause.Root(err), &errMySQL) {
		switch errMySQL.Number {
		case 1062:
			return true
		}
		return false
	}
	var errPostgresql *postgresql.PgError
	if errors.As(cause.Root(err), &errPostgresql) {
		switch errPostgresql.Code {
		case "23505":
			return true
		}
		return false
	}
	return false
}

func IsNoRow(err error) bool {
	return nil != err && dao.IsNoRows(cause.Root(err)) || cause.Match(err, cause.NotFound)
}
