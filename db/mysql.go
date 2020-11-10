package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MySql struct {
	Base
}

func (d *MySql) Init() error {
	DNS := fmt.Sprintf("%s:%s@/%s", d.config.Username, d.config.Password, d.config.Database)

	db, err := sql.Open("mysql", DNS)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(10)

	d.conn = db

	return nil
}
