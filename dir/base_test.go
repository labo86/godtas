package dir

import (
	"strings"
	"testing"
)

func Test_Tmp(t *testing.T) {
	config := Config{
		Type: "tmp",
	}

	dir, err := config.Init()
	if err != nil {
		t.Errorf("%v", err)
	}

	if err != dir.Init() {
		t.Errorf("%v", err)
	}

	filename := dir.Filename("hola")
	if !strings.HasPrefix(filename, "/tmp") {
		t.Errorf("%q : %+v", filename, config)
	}

	if !strings.HasSuffix(filename, "/hola") {
		t.Errorf("%q : %+v", filename, config)
	}
}
