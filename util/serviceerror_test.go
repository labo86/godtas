package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLogError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	w := httptest.NewRecorder()

	errorStr := "SOME ERROR"
	ok := LogError(w, errors.New(errorStr))
	if !ok {
		t.Errorf("log error debe ser exitoso")
		return
	}

	content, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("can't read body: %v", err)
		return
	}

	response := strings.TrimSpace(string(content))
	if got, want := response, ":SERVER_ERROR"; !strings.HasSuffix(got, want) {
		t.Errorf("response got %q, want suffix %q", got, want)
	}

	if got, want := buf.String(), fmt.Sprintf("%s]%s\n", response, errorStr); !strings.HasSuffix(got, want) {
		t.Errorf("response got %q, want %q", got, want)
	}

}

func TestLogServiceError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	w := httptest.NewRecorder()

	errorStr := "SOME ERROR"
	extErrorStr := "EXT_ERROR"
	ok := LogError(w, NewServiceError(extErrorStr, http.StatusBadRequest, errors.New(errorStr)))
	if !ok {
		t.Errorf("log error debe ser exitoso")
		return
	}

	content, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Errorf("can't read body: %v", err)
		return
	}

	response := strings.TrimSpace(string(content))
	if got, want := response, fmt.Sprintf(":%s", extErrorStr); !strings.HasSuffix(got, want) {
		t.Errorf("response got %q, want suffix %q", got, want)
	}

	if got, want := buf.String(), fmt.Sprintf("%s]%s\n", response, errorStr); !strings.HasSuffix(got, want) {
		t.Errorf("response got %q, want %q", got, want)
	}

}

func TestLogErrorNil(t *testing.T) {

	w := httptest.NewRecorder()

	ok := LogError(w, nil)
	if ok {
		t.Errorf("log error debe retornar falso porque el error es nulo")
		return
	}

}
