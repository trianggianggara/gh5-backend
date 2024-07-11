package ctxval

import (
	"context"

	abstraction "gh5-backend/internal/model/base"
)

type key string

var (
	keyAuth key = "x-auth"
	keyTrx  key = "x-trx"
)

func SetAuthValue(ctx context.Context, payload *abstraction.AuthContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetAuthValue(ctx context.Context) *abstraction.AuthContext {
	return ctx.Value(keyAuth).(*abstraction.AuthContext)
}

func SetTrxValue(ctx context.Context, payload *abstraction.TrxContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetTrxValue(ctx context.Context) *abstraction.TrxContext {
	val := ctx.Value(keyTrx)
	if val == nil {
		return nil
	}

	trxCtx, ok := val.(*abstraction.TrxContext)
	if !ok {
		return nil
	}

	return trxCtx
}
