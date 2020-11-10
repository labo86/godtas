package dir

type Config struct {
	Type   string `yaml:"type"`
	Folder string `yaml:"folder"`
}

func (c *Config) Init() (Interface, error) {
	var i Interface
	switch c.Type {
	case "folder":
		i = &Folder{
			Base{
				config: c,
			},
		}
	case "tmp":
		i = &Tmp{
			Base{
				config: c,
			},
		}
	}
	if err := i.Init(); err != nil {
		return nil, err
	}
	return i, nil
}
