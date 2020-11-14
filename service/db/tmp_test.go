package db

import (
	"testing"
)

func Test_Tmp(t *testing.T) {
	d, err := NewTmp()
	if err != nil {
		t.Errorf("deberia crearse : %v", err)
		return
	}

	var value int
	if err := d.Conn().QueryRow(`SELECT 123 as a`).Scan(&value); err != nil {
		t.Errorf("query deberia funcionar: %v", err)
		return
	}

	if got, want := value, 123; got != want {
		t.Errorf("resultado de query : got %q, want %q", got, want)
		return
	}

	if err := d.Close(); err != nil {
		t.Errorf("deberia cerrarse %v", err)
		return
	}
}
