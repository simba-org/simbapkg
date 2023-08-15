package postgres

import (
	"database/sql"
	"gorm.io/gorm"
	"log"
	"time"

	configs "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"

	"golang.org/x/exp/slog"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type DBConnString string

type postgres struct {
	connAttempts int
	connTimeout  time.Duration

	db *sql.DB
}

var _ DBEngine = (*postgres)(nil)

func NewPostgresDB(url configs.DBConnString) (DBEngine, error) {
	slog.Info("CONN", "connect string", url)

	pg := &postgres{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.db, err = sql.Open("postgres", string(url))
		if err != nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	slog.Info("📰 connected to postgresdb 🎉")

	return pg, nil
}

func (p *postgres) Configure(opts ...Option) DBEngine {
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *postgres) GetDB() *sql.DB {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *postgres) GetDbName() string {
	return "postgres"
}

func (p *postgres) GetGormDB() *gorm.DB {
	return nil
}
