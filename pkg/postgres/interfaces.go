package postgres

import "github.com/Bifang-Bird/simbapkg/pkg"

type DBEngine interface {
	pkg.DB
	Configure(...Option) DBEngine
}
