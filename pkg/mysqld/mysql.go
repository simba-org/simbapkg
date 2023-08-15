package mysqld

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
	config "simbapkg/pkg/dbFactory"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type mysqldb struct {
	connAttempts int
	connTimeout  time.Duration
	maxIdleConns int
	maxOpenConns int
	db           *sql.DB
	gormDB       *gorm.DB
}

var _ DBEngine = (*mysqldb)(nil)

func NewMysqlDb(mysqlCfg config.Mysql) (DBEngine, error) {
	slog.Info("CONN", "connect string", mysqlCfg.URL)
	pg := &mysqldb{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		maxOpenConns: mysqlCfg.MaxOpenConns,
		maxIdleConns: mysqlCfg.MaxIdleConns,
	}
	var _db *gorm.DB
	var err error
	for pg.connAttempts > 0 {
		slog.Info(string(mysqlCfg.URL))
		//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
		_db, err = gorm.Open(mysql.Open(string(mysqlCfg.URL)), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		//连接数据库报错，重新连接
		if err != nil {
			log.Printf("mysql is trying to connect, attempts left: %d", pg.connAttempts)
			time.Sleep(pg.connTimeout)
			pg.connAttempts--
			continue
		} else {
			break
		}
	}

	sqlDB, _ := _db.DB()
	if err := sqlDB.Ping(); err != nil {
		slog.Info("mysql connect error:", err.Error())
	} else {
		//设置数据库连接池参数
		sqlDB.SetConnMaxLifetime(time.Minute)
		sqlDB.SetMaxOpenConns(pg.maxOpenConns) //设置数据库连接池最大连接数
		sqlDB.SetMaxIdleConns(pg.maxIdleConns) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
		pg.db = sqlDB
		slog.Info("gorm连接数据库完成")
		pg.gormDB = _db
		slog.Info("📰 connected to mysql 🎉")
	}
	return pg, nil
}

func (p *mysqldb) Configure(opts ...Options) DBEngine {
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *mysqldb) GetDB() *sql.DB {
	return p.db
}

func (p *mysqldb) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *mysqldb) GetDbName() string {
	return "mysql"
}

func (p *mysqldb) GetGormDB() *gorm.DB {
	return p.gormDB
}
