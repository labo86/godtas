package auth0

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Type     string `yaml:"type"`
	Audience string `yaml:"audience"`
	Issuer   string `yaml:"issuer"`
	Url      string `yaml:"url"`
	Kid      string `yaml:"kid"`
	X5c      string `yaml:"x5c"`
}

func (c *Config) Init() (Interface, error) {
	var i Interface
	switch c.Type {
	case "local":
		i = &Local{
			Base{
				config: c,
			},
		}
	case "remote":
		i = &Remote{
			Base{
				config: c,
			},
		}
	default:
		return nil, fmt.Errorf("invalid type %q", c.Type)
	}
	if err := i.Init(); err != nil {
		return nil, fmt.Errorf("auth0.Interface.Init() : %v", err)
	}
	return i, nil
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("auth0.LoadConfig(%q) : ioutil.ReadFile(%q) : %v", filename, filename, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("auth0.LoadConfig(%q) : yaml.Unmarshal() : %v", filename, err)
	}

	return &config, nil
}
