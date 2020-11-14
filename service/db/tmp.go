package db

type Tmp struct {
	Sqlite
}

func NewTmp() *Tmp {
	d := new(Tmp)
	d.Config = &Config{
		Type:     "sqlite3",
		Filename: "testing.db",
		Memory:   true,
	}

	return d
}
