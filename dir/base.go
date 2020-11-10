package dir

import "path/filepath"

type Base struct {
	config   *Config
	basePath string
}

func (d *Base) Filename(filename string) string {
	return filepath.Join(d.basePath, filename)
}
