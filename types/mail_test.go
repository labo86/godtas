package types

import (
	"testing"
)

func TestMail_Validate(t *testing.T) {

	invalidMails := []string{
		"hola",
		"hola@",
		"hola@como@te.com",
		"@hola.com",
	}

	validMails := []string{
		"hola@mail.com",
	}

	for _, value := range invalidMails {
		o := Mail(value)
		if err := o.Validate(); err == nil {
			t.Errorf("deberia ser un correo invalido %q", o)
			return
		}
	}

	for _, value := range validMails {
		o := Mail(value)
		if err := o.Validate(); err != nil {
			t.Errorf("deberia ser un correo valido %q : %v", o, err)
			return
		}
	}

}

func TestMail_Database(t *testing.T) {

	expected := Mail("a@a.com")
	actual := new(Mail)
	if err := AssertInsertSelect(&expected, actual); err != nil {
		t.Error(err)
		return
	}
}

func TestMail_MarshalingJSON(t *testing.T) {

	expected := Mail("a@a.com")
	actual := new(Mail)
	if err := AssertMarshalingJSON(expected, actual); err != nil {
		t.Error(err)
		return
	}
}
