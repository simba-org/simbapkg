package mysqld

import "simbapkg/pkg"

type DBEngine interface {
	pkg.DB
	Configure(...Options) DBEngine
}
