package db

import (
	"io/ioutil"
	"os"
	"testing"
)

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

	db, err := config.Open()
	if err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	if err := db.Open(); err != nil {
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

func Test_Test(t *testing.T) {
	config := Config{
		Type:     "sqlite3",
		Filename: "test.db",
		Memory:   true,
	}

	db, err := config.Open()
	if err != nil {
		t.Errorf("no deberia fallar: %v", err)
	}

	if err := db.Open(); err != nil {
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
		Type:     "sqlite3",
		Filename: "test.db",
		Memory:   true,
	}

	{
		db, err := config.Open()
		if err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Open(); err != nil {
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
		db, err := config.Open()
		if err != nil {
			t.Errorf("no deberia fallar: %v", err)
		}

		if err := db.Open(); err != nil {
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
