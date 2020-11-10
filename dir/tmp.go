package dir

import (
	"io/ioutil"
	"os"
)

type Tmp struct {
	Base
}

func (d *Tmp) Init() error {
	folder, err := ioutil.TempDir("", "test")
	if err != nil {
		return err
	}
	d.basePath = folder
	return nil
}

func (d *Tmp) Close() error {
	return os.RemoveAll(d.basePath)
}
