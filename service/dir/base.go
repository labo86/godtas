package dir

import (
	"os"
	"path/filepath"
)

type Base struct {
	config   *Config
	basePath string
}

func (d *Base) Filename(filename string) string {
	return filepath.Join(d.basePath, filename)
}

func (d *Base) Prepare(filename string) error {
	path := filepath.Dir(d.Filename(filename))

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	return nil
}
