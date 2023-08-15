/*
*

	@author: junwang
	@since: 2023/8/15
	@desc: //TODO

*
*/
package dbFactory

import configs "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"

type Config struct {
	DataSource `yaml:"datasource"`
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
