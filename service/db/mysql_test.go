package db

import (
	"testing"
	"time"
)

func Test_MySql(t *testing.T) {
	config := Config{
		Type:     "mysql",
		Database: "phpunit_test_db",
		Username: "phpunit_test_user",
		Password: "phpunit_test_password",
	}

	db, err := config.Open()
	if err != nil {
		t.Errorf("%+v.Open() got %v", config, err)
	}

	if err := db.Open(); err != nil {
		t.Errorf("%+v.Open() got %v", db, err)
	}

	var value int
	query := `SELECT 123 as a`
	if err := db.Conn().QueryRow(query).Scan(&value); err != nil {
		t.Errorf("%+v: query : %q  got %v", db, query, err)
	}

	if want := 123; value != want {
		t.Errorf("got %d, want %d", value, want)
	}

	if err := db.Close(); err != nil {
		t.Errorf("no deberia fallar el cierre %v", err)
	}
}

func Test_MySqlDate(t *testing.T) {
	config := Config{
		Type:     "mysql",
		Database: "phpunit_test_db",
		Username: "phpunit_test_user",
		Password: "phpunit_test_password",
	}

	db, err := config.Open()
	if err != nil {
		t.Errorf("%+v.Open() got %v", config, err)
	}

	if err := db.Open(); err != nil {
		t.Errorf("%+v.Open() got %v", db, err)
	}

	{
		query := `CREATE TABLE test (id INTEGER, date DATETIME DEFAULT now())`

		if _, err := db.Conn().Exec(query); err != nil {
			t.Errorf("%+v: query : %q  got %v", db, query, err)
		}
	}

	{
		query := `INSERT INTO test (id) VALUES(?)`

		if _, err := db.Conn().Exec(query, 1); err != nil {
			t.Errorf("%+v: query : %q  got %v", db, query, err)
		}
	}

	{
		var value time.Time
		query := `SELECT date FROM test LIMIT 1`
		if err := db.Conn().QueryRow(query).Scan(&value); err != nil {
			t.Errorf("%+v: query : %q  got %v", db, query, err)
		}
	}

	if err := db.Close(); err != nil {
		t.Errorf("no deberia fallar el cierre %v", err)
	}
}
