package db

type Tmp struct {
	Sqlite
}

func NewTmp() *Tmp {
	d := new(Tmp)
	d.config = &Config{
		Type:     "sqlite3",
		Filename: "test.db",
		Memory:   true,
	}

	return d
}
