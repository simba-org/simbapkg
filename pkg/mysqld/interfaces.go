package mysqld

import "github.com/Bifang-Bird/simbapkg/pkg"

type DBEngine interface {
	pkg.DB
	Configure(...Options) DBEngine
}
