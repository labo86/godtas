package db

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_Test(t *testing.T) {
	config := Config{
		Type: "test",
	}

	db, err := config.Init()
	if err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	if err := db.Init(); err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	var value int
	if err != db.Conn().QueryRow(`SELECT 123 as a`).Scan(&value) {
		t.Errorf("no deberia fallar: %v", err)
	}

	if value != 123 {
		t.Errorf("el valor deberia ser 123: %q", value)
	}

	if err := db.Close(); err != nil {
		t.Errorf("no deberia fallar el cierre %v", err)
	}
}

func Test_CreateTableDifferentConnectionCache(t *testing.T) {
	config := Config{
		Type: "test",
	}

	{
		db, err := config.Init()
		if err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Init(); err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if _, err := db.Conn().Exec(`CREATE TABLE test (id TEXT)`); err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Close(); err != nil {
			t.Errorf("no deberia fallar el cierre %v", err)
		}

	}

	{
		db, err := config.Init()
		if err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Init(); err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if _, err := db.Conn().Exec(`CREATE TABLE test (id TEXT)`); err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Close(); err != nil {
			t.Errorf("no deberia fallar el cierre %v", err)
		}
	}
}

func Test_Sqlite3(t *testing.T) {
	filename, err := ioutil.TempFile("", "temp_db")
	if err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	defer os.Remove(filename.Name())

	config := Config{
		Type:     "sqlite3",
		Filename: filename.Name(),
		Memory:   false,
	}

	db, err := config.Init()
	if err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	if err := db.Init(); err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	var value int
	if err != db.Conn().QueryRow(`SELECT 123 as a`).Scan(&value) {
		t.Errorf("no deberia fallar: %v", err)
	}

	if value != 123 {
		t.Errorf("el valor deberia ser 123: %q", value)
	}

	if err := db.Close(); err != nil {
		t.Errorf("no deberia fallar el cierre %v", err)
	}
}

func Test_MySql(t *testing.T) {
	config := Config{
		Type:     "mysql",
		Database: "phpunit_test_db",
		Username: "phpunit_test_user",
		Password: "phpunit_test_password",
	}

	db, err := config.Init()
	if err != nil {
		t.Errorf("%+v.Init() got %v", config, err)
	}

	if err := db.Init(); err != nil {
		t.Errorf("%+v.Init() got %v", db, err)
	}

	var value int
	if query := `SELECT 123 as a`; err != db.Conn().QueryRow(query).Scan(&value) {
		t.Errorf("%+v: query : %q  got %v", db, query, err)
	}

	if want := 123; value != want {
		t.Errorf("got %d, want %d", value, want)
	}

	if err := db.Close(); err != nil {
		t.Errorf("no deberia fallar el cierre %v", err)
	}
}
