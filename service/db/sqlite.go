package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Base
}

func (d *Sqlite) Open() error {
	var DNS string
	if d.Config.Memory {
		DNS = fmt.Sprintf("file:%s?cache=shared&mode=memory", d.Config.Filename)
	} else {
		DNS = fmt.Sprintf("file:%s?cache=shared", d.Config.Filename)
	}

	db, err := sql.Open("sqlite3", DNS)
	if err != nil {
		return fmt.Errorf("can't open db : %v", err)
	}

	d.conn = db
	return nil
}
