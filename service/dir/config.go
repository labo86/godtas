package dir

import "fmt"

type Config struct {
	Type   string `yaml:"type"`
	Folder string `yaml:"folder"`
}

func (c *Config) Open() (Dir, error) {
	var i Dir
	switch c.Type {
	case "folder":
		i = &Folder{
			Base{
				config: c,
			},
		}
	default:
		return nil, fmt.Errorf("wrong type %q", c.Type)
	}
	if err := i.Open(); err != nil {
		return nil, err
	}
	return i, nil
}
