package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const MysqlDateFormat = "2006-01-02 15:04:05"

type Timestamp struct {
	time.Time
}

func (o Timestamp) Value() (driver.Value, error) {

	return o.Format(MysqlDateFormat), nil

}

func (o *Timestamp) Scan(value interface{}) error {

	switch value.(type) {
	case string:
		newTime, err := time.Parse(MysqlDateFormat, value.(string))
		if err != nil {
			return fmt.Errorf("%q is not a valid date: %v", value.(string), err)
		}

		o.Time = newTime
	case time.Time:
		o.Time = value.(time.Time)
	default:
		return fmt.Errorf("not a valid date type : %v", value)
	}

	return nil
}
