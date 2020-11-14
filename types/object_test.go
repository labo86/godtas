package types

import (
	"fmt"
	"testing"
)

func TestObject_Database(t *testing.T) {

	expected := &Object{
		"name":    "edwin",
		"surname": "rodriguez",
		"list":    []string{"1", "2", "3"},
	}
	actual := &Object{}
	if err := AssertInsertSelect(expected, actual); err != nil {
		if fmt.Sprint(expected) != fmt.Sprint(actual) {
			t.Error(err)
			return
		}
	}
}

func TestObject_JSON(t *testing.T) {

	expected := &Object{
		"name":    "edwin",
		"surname": "rodriguez",
		"list":    []string{"1", "2", "3"},
	}
	actual := &Object{}
	if err := AssertMarshalingJSON(expected, actual); err != nil {

		if fmt.Sprint(expected) != fmt.Sprint(actual) {
			t.Error(err)
			return
		}
	}
}
