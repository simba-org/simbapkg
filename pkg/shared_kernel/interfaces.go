package sharedkernel

import (
	"context"
	"time"
)

type (
	Handler     func(context.Context, DomainEvent) (interface{}, error)
	Handlers    []Handler
	DomainEvent interface {
		CreateAt() time.Time
		Identity() string
		InitHandler(...Handler)                          // 初始化事件处理函数
		Handle(ctx context.Context) (interface{}, error) // 调用事件处理函数
		IgnoreErr() bool                                 // 是否忽略事件处理错误
	}
)
