package util

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestEncodeAsJson(t *testing.T) {
	w := httptest.NewRecorder()

	slice := make([]string, 0)

	EncodeAsJson(w, slice)

	content, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("can't read body: %v", err)
		return
	}

	if got, want := string(content), "[]\n"; got != want {
		t.Errorf("response got %q, want suffix %q", got, want)

	}
}
