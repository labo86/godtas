package auth0

import jwtmiddleware "github.com/auth0/go-jwt-middleware"

type Interface interface {
	Init() error
	Close() error
	MiddleWare() *jwtmiddleware.JWTMiddleware
}
