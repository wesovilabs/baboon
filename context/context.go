package context

import "context"

type RootContext interface {
}
type rootContext struct {
}

func New(ctx context.Context) RootContext {
	return &rootContext{}
}
