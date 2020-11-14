package types

import (
	"fmt"
	"strings"
)

type Mail string

func (o Mail) Validate() error {

	str := string(o)
	values := strings.Split(str, "@")
	if len(values) != 2 {
		return fmt.Errorf("have more than one @ character")
	}

	user := values[0]
	domain := values[1]

	if len(user) <= 0 {
		return fmt.Errorf("empty user")
	}

	if len(domain) <= 0 {
		return fmt.Errorf("empty domain")
	}

	return nil
}
