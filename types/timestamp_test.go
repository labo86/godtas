package types

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestTimestamp_Database(t *testing.T) {

	expected := Timestamp{
		time.Now().UTC().Truncate(time.Second),
	}
	actual := Timestamp{}
	if err := AssertInsertSelect(&expected, &actual); err != nil {
		t.Error(err)
		return
	}
}

func TestTimestamp_DatabaseDatetime(t *testing.T) {

	expected := Timestamp{
		time.Now().UTC().Truncate(time.Second),
	}
	actual := Timestamp{}

	d, err := OpenDBTmp(`CREATE TABLE types (id TEXT, value DATETIME)`)
	if err != nil {
		t.Errorf("can't open tmp db : %v", err)
		return
	}
	defer d.Close()

	err = func() error {
		if err := Insert(d, "1", &expected); err != nil {
			return err
		}

		if err := Select(d, "1", &actual); err != nil {
			return err
		}

		if !reflect.DeepEqual(expected, actual) {
			return fmt.Errorf("got %v , want %v", actual, expected)
		}

		return nil
	}()

	if err != nil {
		t.Error(err)
		return
	}
}

func TestTimestamp_JSON(t *testing.T) {

	expected := Timestamp{
		time.Now().UTC().Truncate(time.Second),
	}
	actual := Timestamp{}
	if err := AssertMarshalingJSON(&expected, &actual); err != nil {
		t.Error(err)
		return
	}
}
