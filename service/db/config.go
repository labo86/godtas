package db

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Type     string `yaml:"type"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	//for sqlite
	Filename string `yaml:"filename"`
	//for sqlite
	Memory bool `yaml:"memory"`
}

func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read file %q : %v", filename, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall yaml %q : %v", filename, err)
	}

	return &config, nil
}

func (c *Config) Open() (DB, error) {
	var db DB
	switch c.Type {
	case "sqlite3":
		db = &Sqlite{
			Base{
				config: c,
			},
		}
	case "mysql":
		db = &MySql{
			Base{
				config: c,
			},
		}
	default:
		return nil, fmt.Errorf("wrong type %q", c.Type)
	}
	if err := db.Open(); err != nil {
		return nil, err
	}
	return db, nil
}
