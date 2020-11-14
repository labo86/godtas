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

func NowTimestamp() *Timestamp {
	t := &Timestamp{
		time.Now().UTC().Truncate(time.Duration(time.Second)),
	}
	return t
}

func (o Timestamp) Value() (driver.Value, error) {

	return o.Format(MysqlDateFormat), nil

}

func (o *Timestamp) Scan(value interface{}) error {

	concreteValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("converted value is not a string : %+v", value)
	}

	newTime, err := time.Parse(MysqlDateFormat, concreteValue)
	if err != nil {
		return fmt.Errorf("%q is not a valid date: %v", concreteValue, err)
	}

	o.Time = newTime

	return nil
}
