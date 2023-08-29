package pkg

import (
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/rabbitmq/publisher"
	"context"
	"database/sql"
	config "github.com/Bifang-Bird/simbapkg/pkg/dbconfig"
	"gorm.io/gorm"
)

type DB interface {
	GetDbName() string
	GetDB() *sql.DB
	Close()
	GetGormDB() *gorm.DB
}

type MqPublisher interface {
	Configure(...publisher.Option)
	DelayConfigure(...publisher.DeplayOption)
	Publish(context.Context, []byte, string, config.MqConfig) error
	DelayPublish(context.Context, []byte, string, int64) error
}
