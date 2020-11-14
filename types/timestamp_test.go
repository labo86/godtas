package types

import (
	"testing"
)

func TestTimestamp_Database(t *testing.T) {

	expected := NowTimestamp()
	actual := &Timestamp{}
	if err := AssertInsertSelect(expected, actual); err != nil {
		t.Error(err)
		return
	}
}

func TestTimestamp_JSON(t *testing.T) {

	expected := NowTimestamp()
	actual := &Timestamp{}
	if err := AssertMarshalingJSON(expected, actual); err != nil {
		t.Error(err)
		return
	}
}
