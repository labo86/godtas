package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Object map[string]interface{}

func (o Object) Value() (driver.Value, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("can't marshal value %+v: %v", o, err)
	}

	return data, nil
}

func (o *Object) Scan(value interface{}) error {

	concreteValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("converted value is not a []byte : %+v", value)
	}

	if err := json.Unmarshal(concreteValue, o); err != nil {
		return fmt.Errorf("%q is not a valid json", string(concreteValue))
	}

	return nil
}
