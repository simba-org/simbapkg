package mysqld

import "time"

type Options func(*mysqldb)

func ConnAttempts(attempts int) Options {
	return func(p *mysqldb) {
		p.connAttempts = attempts
	}
}

func ConnTimeout(timeout time.Duration) Options {
	return func(p *mysqldb) {
		p.connTimeout = timeout
	}
}
