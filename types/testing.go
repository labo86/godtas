package types

import (
	"encoding/json"
	"fmt"
	"github.com/labo86/godtas/service/db"
	"reflect"
)

type DBTypeTmp struct {
	db.Sqlite
}

func OpenDBTmp() (*DBTypeTmp, error) {
	o := new(DBTypeTmp)
	o.Config = &db.Config{
		Type:     "sqlite3",
		Filename: "testing.db",
		Memory:   true,
	}

	if err := o.Open(); err != nil {
		return nil, fmt.Errorf("can't open tmp db: %v", err)
	}

	{
		d := o.Conn()
		if _, err := d.Exec(`CREATE TABLE types (id TEXT, value TEXT)`); err != nil {
			_ = d.Close()
			return nil, fmt.Errorf("can't create tmp table : %v", err)
		}
	}

	return o, nil

}

func (o *DBTypeTmp) Insert(id interface{}, value interface{}) error {
	d := o.Conn()
	if _, err := d.Exec(`INSERT INTO types (id, value) VALUES (?, ?)`, id, value); err != nil {
		return fmt.Errorf("can't insert type %q : %v", id, err)
	}
	return nil
}

func (o *DBTypeTmp) Select(id interface{}, value interface{}) error {
	d := o.Conn()
	if err := d.QueryRow(`SELECT value FROM types WHERE id = ?`, id).Scan(value); err != nil {
		return fmt.Errorf("can't retrieve type %q : %v", id, err)
	}
	return nil
}

func AssertInsertSelect(expected interface{}, actual interface{}) error {
	d, err := OpenDBTmp()
	if err != nil {
		return fmt.Errorf("can't open tmp db : %v", err)
	}
	defer d.Close()

	if err := d.Insert("1", expected); err != nil {
		return err
	}

	if err := d.Select("1", actual); err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("got %v , want %v", actual, expected)
	}

	return nil
}

func AssertMarshalingJSON(expected interface{}, actual interface{}) error {

	data, err := json.Marshal(expected)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, actual); err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("got %v , want %v", actual, expected)
	}

	return nil
}
