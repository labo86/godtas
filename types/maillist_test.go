package types

import (
	"reflect"
	"testing"
)

func TestMailList_Normalized(t *testing.T) {
	o := MailList{"c@c.com", "b@b.com", "a@a.com", "b@b.com"}

	r := o.Normalized()

	expected := MailList{"a@a.com", "b@b.com", "c@c.com"}
	if got, want := r, expected; !reflect.DeepEqual(got, want) {
		t.Errorf("normalize : got %v, want %v", got, want)
	}
}

func TestMailList_Database(t *testing.T) {

	expected := &MailList{"a@a.com", "b@b.com", "c@c.com"}
	actual := &MailList{}
	if err := AssertInsertSelect(expected, actual); err != nil {
		t.Error(err)
		return
	}
}

func TestMailList_MarshalingJSON(t *testing.T) {

	expected := &MailList{"a@a.com", "b@b.com", "c@c.com"}
	actual := &MailList{}
	if err := AssertMarshalingJSON(expected, actual); err != nil {
		t.Error(err)
		return
	}
}
