package generics

import (
	"context"
)

func FromContext[T any](ctx context.Context) T {
	return Cast[T](ctx.Value(TypeOf[T]()))
}

func WithValue[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, TypeOf[T](), v)
}
