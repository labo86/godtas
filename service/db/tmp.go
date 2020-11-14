package db

type Tmp struct {
	Sqlite
}

func NewTmp() (*Tmp, error) {
	d := new(Tmp)
	d.Config = &Config{
		Type:     "sqlite3",
		Filename: "testing.db",
		Memory:   true,
	}

	err := d.Open()
	return d, err
}
