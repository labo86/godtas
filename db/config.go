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

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading config file %q : %v", data, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Init() (Interface, error) {
	var db Interface
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
		db = &Sqlite{
			Base{
				config: &Config{
					Type:     "sqlite3",
					Filename: "test.db",
					Memory:   true,
				},
			},
		}
	}
	if err := db.Init(); err != nil {
		return nil, err
	}
	return db, nil
}
