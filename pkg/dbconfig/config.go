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
}

type DataSource struct {
	Type  string `env-required:"true" yaml:"type" env:"TYPE"`
	Mysql Mysql  `env-required:"true" yaml:"mysql" env:"MYSQL"`
	PG    PG     `env-required:"true" yaml:"postgres" env:"POSTGRES"`
}

type PG struct {
	PoolMax int                  `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	DsnURL  configs.DBConnString `env-required:"true" yaml:"dsn_url" env:"PG_DSN_URL"`
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
