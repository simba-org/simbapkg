/*
*

	@author: junwang
	@since: 2023/8/15
	@desc: //TODO

*
*/
package dbconfig

import configs "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"

type Config struct {
	DataSource  `yaml:"datasource"`
	LoadBalance `yaml:"loadBalance"`
	RabbitMQ    `yaml:"rabbitmq" ,env-required:"false" ,env:"RABBITMQ"`
}

type (
	RabbitMQ struct {
		URL      string            `yaml:"url" ,env-required:"false" ,env:"RABBITMQ_URL"`
		Publish  []PublishProfile  `yaml:"publish" ,env-required:"false" ,env:"PUBLISH"`
		Consumer []ConsumerProfile `yaml:"consumer" ,env-required:"false" ,env:"CONSUMER"`
	}
	PublishProfile struct {
		Type string      `env-required:"true" yaml:"type" env:"TYPE"`
		Body PublishBody `env-required:"true" yaml:"body" env:"BODY"`
	}

	PublishBody struct {
		ExchangeName    string `env-required:"true" yaml:"exchangeName" env:"EXCHANGE_NAME"`
		BindingKey      string `env-required:"true" yaml:"bindingKey" env:"BINDING_KEY"`
		MessageTypeName string `env-required:"true" yaml:"messageTypeName" env:"MESSAGE_TYPE_NAME"`
	}
	ConsumerProfile struct {
		Type string       `env-required:"true" yaml:"type" env:"TYPE"`
		Body ConsumerBody `env-required:"true" yaml:"body" env:"BODY"`
	}

	ConsumerBody struct {
		ExchangeName string `env-required:"true" yaml:"exchangeName" env:"EXCHANGE_NAME"`
		BindingKey   string `env-required:"true" yaml:"bindingKey" env:"BINDING_KEY"`
		ConsumerTag  string `env-required:"true" yaml:"consumerTag" env:"CONSUMER_TAG"`
		QueueName    string `env-required:"true" yaml:"queueName" env:"QUEUE_NAME"`
	}
)

type DataSource struct {
	Type  string `yaml:"type" ,env-required:"true" ,env:"TYPE"`
	Mysql Mysql  `yaml:"mysql" ,env-required:"true" ,env:"MYSQL"`
	PG    PG     `yaml:"postgres" ,env-required:"false" ,env:"POSTGRES"`
}

type PG struct {
	PoolMax int                  `yaml:"pool_max" ,env-required:"false" ,env:"PG_POOL_MAX"`
	DsnURL  configs.DBConnString `yaml:"dsn_url" ,env-required:"false" ,env:"PG_DSN_URL"`
}

type Mysql struct {
	MaxOpenConns int                  `env-required:"true" yaml:"max_open_conns" env:"MAX_OPEN_CONNS"`
	MaxIdleConns int                  `env-required:"true" yaml:"max_idle_conns" env:"MAX_IDLE_CONNS"`
	URL          configs.DBConnString `env-required:"true" yaml:"url" env:"URL"`
}

type (
	LoadBalance struct {
		Specify    bool   `yaml:"specify" env:"SPECIFY"`
		Channel    string `env-required:"true" yaml:"channel" env:"CHANNEL"`
		SelectMode `env-required:"true" yaml:"selectMode" env:"SELECT_MODE"`
	}
	SelectMode struct {
		Strategy int       `yaml:"strategy" env:"STRATEGY"`
		Weight   []*Weight `env-required:"true" yaml:"weight" env:"WEIGHT"`
	}
	Weight struct {
		Chan  string `env-required:"true" yaml:"chan" env:"CHAN"`
		Value string `env-required:"true" yaml:"value" env:"VALUE"`
	}
)

type MqConfig struct {
	ExchangeName    string
	BindingKey      string
	MessageTypeName string
}
