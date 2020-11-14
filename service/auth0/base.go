package auth0

import (
	"crypto/rsa"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type Base struct {
	certs      map[string]*rsa.PublicKey
	config     *Config
	middleWare *jwtmiddleware.JWTMiddleware
}

func (d *Base) Close() error {
	return nil
}

func (d *Base) MiddleWare() *jwtmiddleware.JWTMiddleware {
	if d.middleWare == nil {
		d.middleWare = jwtmiddleware.New(jwtmiddleware.Options{

			ValidationKeyGetter: d.GetValidationKeyForToken,

			//el nombre de la variable de contexto en donde irá la información de usuario
			UserProperty: "user",

			SigningMethod: jwt.SigningMethodRS256,
		})
	}
	return d.middleWare
}

/*
 Obtiene la llave publica para un kid, el kid debe venir en el header del JWT
 Por lo menos asi viene en auth0
*/
func (d *Base) getPublicKey(token *jwt.Token) (*rsa.PublicKey, error) {
	kid := token.Header["kid"].(string)
	cert, ok := d.certs[kid]

	if !ok {
		return nil, fmt.Errorf("cert key %q not found", kid)
	}

	return cert, nil
}

func (d *Base) GetValidationKeyForToken(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(jwt.MapClaims)

	// Verify 'aud' claim
	audience := d.config.Audience
	checkAud := claims.VerifyAudience(audience, false)
	if !checkAud {
		return nil, fmt.Errorf("audience should be %q instead %+v", audience, claims)
	}

	// Verify 'iss' claim
	issuer := d.config.Issuer
	checkIss := claims.VerifyIssuer(issuer, false)
	if !checkIss {
		return nil, fmt.Errorf("issuer should be %q instead %+v", issuer, claims)
	}

	publicKey, err := d.getPublicKey(token)
	if err != nil {
		return nil, fmt.Errorf("public key not available: %v", err)
	}
	return publicKey, nil
}
