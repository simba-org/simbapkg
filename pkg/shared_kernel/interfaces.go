package sharedkernel

import (
	"context"
	"time"
)

type (
	Handler     func(context.Context, DomainEvent) (interface{}, error)
	DomainEvent interface {
		CreateAt() time.Time
		Identity() string
		InitHandler(Handler)
		Handle(ctx context.Context) (interface{}, error)
	}
)
