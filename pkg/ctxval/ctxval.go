package ctxval

import (
	"context"

	abstraction "gh5-backend/internal/model/base"
)

type key string

var keyAuth key = "x-auth"

func SetAuthValue(ctx context.Context, payload *abstraction.AuthContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetAuthValue(ctx context.Context) *abstraction.AuthContext {
	return ctx.Value(keyAuth).(*abstraction.AuthContext)
}
