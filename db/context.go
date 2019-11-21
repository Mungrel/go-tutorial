package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type dbContextKey int

const dbCtxKey int = iota

func ContextWithDB(ctx context.Context, client *sqlx.DB) context.Context {
	return context.WithValue(ctx, dbCtxKey, client)
}

func ContextDB(ctx context.Context) *sqlx.DB {
	return ctx.Value(dbCtxKey).(*sqlx.DB)
}
