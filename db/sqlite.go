package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Base
}

func (d *Sqlite) Init() error {
	var DNS string
	if d.config.Memory {
		DNS = fmt.Sprintf("file:%s?cache=shared&mode=memory", d.config.Filename)
	} else {
		DNS = fmt.Sprintf("file:%s?cache=shared", d.config.Filename)
	}

	db, err := sql.Open("sqlite3", DNS)
	if err != nil {
		return err
	}

	d.conn = db
	return nil
}
