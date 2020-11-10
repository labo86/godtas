package auth0

import (
	"bytes"
	"encoding/json"
)

type Local struct {
	Base
}

func (d *Local) Init() error {
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
