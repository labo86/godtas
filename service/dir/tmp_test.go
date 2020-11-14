package dir

import (
	"strings"
	"testing"
)

func Test_Tmp(t *testing.T) {
	dir, err := NewTmp()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	filename := dir.Filename("hola")

	if got, want := filename, "/tmp"; !strings.HasPrefix(got, want) {
		t.Errorf("filename %q deberia estar en el dir %q", got, want)
		return
	}

	if got, want := filename, "/hola"; !strings.HasSuffix(got, want) {
		t.Errorf("filename %q deberia ser %q", got, want)
		return
	}
}
