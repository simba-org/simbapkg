package pkg

import (
	"database/sql"
	"gorm.io/gorm"
)

type DB interface {
	GetDbName() string
	GetDB() *sql.DB
	Close()
	GetGormDB() *gorm.DB
}
