package auth0

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Base struct {
	certs  map[string]*rsa.PublicKey
	config *Config
}

func (d *Base) Close() error {
	return nil
}

/*
 Obtiene la llave publica para un kid, el kid debe venir en el header del JWT
 Por lo menos asi viene en auth0
*/
func (d *Base) getPublicKey(token *jwt.Token) (interface{}, error) {
	kid := token.Header["kid"].(string)
	cert, ok := d.certs[kid]

	if !ok {
		return nil, fmt.Errorf("cert key %q not found", kid)
	}

	return cert, nil
}

func (d *Base) CheckJWT(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, d.getPublicKey)
	if err != nil {
		return nil, fmt.Errorf("can't parse token %q : %v", token, err)
	}

	if got, want := parsedToken.Header["alg"], jwt.SigningMethodRS256.Alg(); got != want {
		return nil, fmt.Errorf("wront signing method, want %q, got %q", want, got)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token %q", token)
	}

	return parsedToken, nil
}
