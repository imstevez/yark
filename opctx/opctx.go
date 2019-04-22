package opctx

import "context"

type CtxKey string

func Set(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, CtxKey(key), value)
}

func Get(ctx context.Context, key string) interface{} {
	return ctx.Value(CtxKey(key))
}
