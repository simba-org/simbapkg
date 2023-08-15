package postgres

import "simbapkg/pkg"

type DBEngine interface {
	pkg.DB
	Configure(...Option) DBEngine
}
