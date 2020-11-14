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

func (d *MySql) Open() error {
	DNS := fmt.Sprintf("%s:%s@/%s", d.Config.Username, d.Config.Password, d.Config.Database)

	db, err := sql.Open("mysql", DNS)
	if err != nil {
		return fmt.Errorf("can't open db : %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(10)

	d.conn = db

	return nil
}
