package dbFactory

import (
	"github.com/Bifang-Bird/simbapkg/pkg"
	"github.com/Bifang-Bird/simbapkg/pkg/dbconfig"
	"github.com/Bifang-Bird/simbapkg/pkg/mysqld"
	"github.com/Bifang-Bird/simbapkg/pkg/postgres"
)

func GetDb(ds dbconfig.DataSource) (pkg.DB, error) {
	switch ds.Type {
	case "mysql":
		pg, err := mysqld.NewMysqlDb(ds.Mysql)
		return pg, err
	case "postgres":
		pg, err := postgres.NewPostgresDB(ds.PG.DsnURL)
		return pg, err
	}
	return nil, nil
}
