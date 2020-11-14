package db

import "testing"

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
