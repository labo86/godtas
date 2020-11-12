package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Local struct {
	Base
}

func (d *Local) Init() error {
	if d.config.Kid == "" {
		return fmt.Errorf("auth0.Local.Init() : Kid is empty")
	}

	if d.config.X5c == "" {
		return fmt.Errorf("auth0.Local.Init() : X5c is empty")
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
		return err
	}

	reader := bytes.NewReader(content)
	certs, err := ParseJsonKeys(reader)
	if err != nil {
		return err
	}
	d.certs = certs

	return nil
}
