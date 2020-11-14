package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Local struct {
	Base
}

func (d *Local) Open() error {
	if d.config.Kid == "" {
		return fmt.Errorf("empty Kid")
	}

	if d.config.X5c == "" {
		return fmt.Errorf("empty X5c")
	}

	jwks := JWKS{
		Keys: []JWK{
			{
				Kid: d.config.Kid,
				X5c: []string{
					d.config.X5c,
				},
			},
		},
	}

	content, err := json.Marshal(jwks)
	if err != nil {
		return fmt.Errorf("can't marshal jwks : %v", err)
	}

	reader := bytes.NewReader(content)
	certs, err := ParseJsonKeys(reader)
	if err != nil {
		return fmt.Errorf("can't parse json keys: %v", err)
	}
	d.certs = certs

	return nil
}
