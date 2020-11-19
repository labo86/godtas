package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
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

	dir := path.Dir(d.Config.Filename)

	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return fmt.Errorf("can't create database %q dir : %v", d.Config.Filename, err)
	}

	db, err := sql.Open("sqlite3", DNS)
	if err != nil {
		return fmt.Errorf("can't open db : %v", err)
	}

	d.conn = db
	return nil
}
