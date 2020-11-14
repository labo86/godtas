package auth0

import jwtmiddleware "github.com/auth0/go-jwt-middleware"

type Auth0 interface {
	Open() error
	Close() error
	MiddleWare() *jwtmiddleware.JWTMiddleware
}
