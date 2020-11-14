package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
)

func OpenDBTmp(ddl string) (*sql.DB, error) {

	d, err := sql.Open("sqlite3", "file:%s?cache=shared&mode=memory")
	if err != nil {
		return nil, fmt.Errorf("can't open tmp db: %v", err)
	}

	if _, err := d.Exec(ddl); err != nil {
		_ = d.Close()
		return nil, fmt.Errorf("can't create tmp table : %v", err)
	}

	return d, nil
}

func OpenDBTypeTmp() (*sql.DB, error) {
	ddl := `CREATE TABLE types (id TEXT, value TEXT)`
	return OpenDBTmp(ddl)
}

func Insert(o *sql.DB, id interface{}, value interface{}) error {
	if _, err := o.Exec(`INSERT INTO types (id, value) VALUES (?, ?)`, id, value); err != nil {
		return fmt.Errorf("can't insert type %q : %v", id, err)
	}
	return nil
}

func Select(o *sql.DB, id interface{}, value interface{}) error {

	if err := o.QueryRow(`SELECT value FROM types WHERE id = ?`, id).Scan(value); err != nil {
		return fmt.Errorf("can't retrieve type %q : %v", id, err)
	}
	return nil
}

func AssertInsertSelect(expected interface{}, actual interface{}) error {
	d, err := OpenDBTypeTmp()
	if err != nil {
		return fmt.Errorf("can't open tmp db : %v", err)
	}
	defer d.Close()

	if err := Insert(d, "1", expected); err != nil {
		return err
	}

	if err := Select(d, "1", actual); err != nil {
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
