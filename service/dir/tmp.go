package dir

import (
	"fmt"
	"io/ioutil"
)

type Tmp struct {
	Folder
}

func NewTmp() (*Tmp, error) {
	folder, err := ioutil.TempDir("", "testing")
	if err != nil {
		return nil, fmt.Errorf("can't create temp dir %v", err)
	}

	d := new(Tmp)
	d.config = &Config{
		Type:   "folder",
		Folder: folder,
	}

	err = d.Open()
	return d, err
}
