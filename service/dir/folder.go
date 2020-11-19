package dir

import (
	"os"
)

type Folder struct {
	Base
}

func (d *Folder) Open() error {
	folder := d.config.Folder

	if err := os.MkdirAll(folder, 0644); err != nil {
		return err
	}
	d.basePath = folder
	return nil
}

func (d *Folder) Close() error {
	return nil
}
