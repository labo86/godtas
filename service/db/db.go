package db

import "database/sql"

type DB interface {
	Open() error
	Close() error
	Conn() *sql.DB
}
