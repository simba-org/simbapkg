package publish

import (
	"context"
	"github.com/Bifang-Bird/simbapkg/pkg"
	"github.com/Bifang-Bird/simbapkg/pkg/dbconfig"

	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/publisher"

	"github.com/google/wire"
)

var (
	PublisherSet = wire.NewSet(NewMqPublisher)
)

type (
	MqPublisher struct {
		pub      publisher.EventPublisher
		delayPub publisher.EventDelayPublisher
	}
)

func NewMqPublisher(pub publisher.EventPublisher, delayPub publisher.EventDelayPublisher) pkg.MqPublisher {
	return &MqPublisher{
		pub:      pub,
		delayPub: delayPub,
	}
}

func (p *MqPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *MqPublisher) Publish(ctx context.Context, body []byte, contentType string, rabbitmqCFG dbconfig.MqConfig) error {
	p.Configure(
		publisher.ExchangeName(rabbitmqCFG.ExchangeName),
		publisher.BindingKey(rabbitmqCFG.BindingKey),
		publisher.MessageTypeName(rabbitmqCFG.MessageTypeName),
	)
	return p.pub.Publish(ctx, body, contentType)
}

func (p *MqPublisher) DelayConfigure(opts ...publisher.DeplayOption) {
	p.delayPub.Configure(opts...)
}

func (p *MqPublisher) DelayPublish(ctx context.Context, body []byte, contentType string, delay int64) error {
	return p.delayPub.DelayPublish(ctx, body, contentType, delay)
}
