package db

import "database/sql"

type Base struct {
	config *Config
	conn   *sql.DB
}

func (d *Base) Conn() *sql.DB {
	return d.conn
}

func (d *Base) Close() error {
	if d.conn != nil {
		err := d.conn.Close()
		d.conn = nil
		return err
	}
	return nil
}
