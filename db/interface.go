package db

import "database/sql"

type Interface interface {
	Init() error
	Close() error
	Conn() *sql.DB
}
