package auth0

import "github.com/dgrijalva/jwt-go"

type Auth0 interface {
	Open() error
	Close() error
	CheckJWT(string) (*jwt.Token, error)
}
