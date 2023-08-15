package dbFactory

import (
	"simbapkg/pkg"
	"simbapkg/pkg/dbconfig"

	"simbapkg/pkg/mysqld"

	"simbapkg/pkg/postgres"
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
