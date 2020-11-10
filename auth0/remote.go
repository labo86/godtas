package auth0

import (
	"fmt"
	"net/http"
)

type Remote struct {
	Base
}

func (d *Remote) Init() error {

	url := d.config.Url

	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("can't download auth0 pem cert %q : %v", url, err)
	}

	certs, err := ParseJsonKeys(resp.Body)
	if err != nil {
		return fmt.Errorf("error parsing cert content: %v", err)
	}

	d.certs = certs
	return nil
}
