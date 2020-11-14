package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
)

type MailList []Mail

func (o MailList) Value() (driver.Value, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("can't marshal value %+v: %v", o, err)
	}

	return data, nil
}

func (o *MailList) Scan(value interface{}) error {

	concreteValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("converted value is not a []byte : %+v", value)
	}

	if err := json.Unmarshal(concreteValue, o); err != nil {
		return fmt.Errorf("%q is not a valid json: %v", string(concreteValue), err)
	}

	return nil
}

func (o MailList) Normalized() MailList {

	keys := make(map[Mail]bool)

	for _, value := range o {
		keys[value] = true
	}

	r := MailList{}

	for k := range keys {
		r = append(r, k)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return r

}
