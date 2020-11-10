package auth0

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"net/http"
)

type JWK struct {
	Kid string   `json:"kid"`
	X5c []string `json:"x5c"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

func RequestUser(r *http.Request) (string, error) {
	user := r.Context().Value("user")

	if user == nil {
		return "", fmt.Errorf("auth0: no se encontró la authorization en user")
	}

	token := user.(*jwt.Token)
	if token == nil {
		return "", fmt.Errorf("auth0: la variable no es un jwt : %+v", token)
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims == nil {
		return "", fmt.Errorf("auth0: el jwd no tiene claims : %+v", claims)
	}

	sub := claims["sub"].(string)
	if sub == "" {
		return "", fmt.Errorf("auth0: claim no tiene variable valor sub o es vacio : %+v", claims)
	}

	return sub, nil
}

/*
En el formato que viene en jwks.json de auth0
@see https://auth0.com/docs/tokens/json-web-tokens/json-web-key-sets
*/
func ParseJsonKeys(jsonData io.Reader) (map[string]*rsa.PublicKey, error) {
	var jwks = JWKS{}
	err := json.NewDecoder(jsonData).Decode(&jwks)
	if err != nil {
		return nil, fmt.Errorf("can't decode response body :%v", err)
	}

	certs := map[string]*rsa.PublicKey{}

	for _, v := range jwks.Keys {
		pem := fmt.Sprintf(`-----BEGIN CERTIFICATE-----
%s
-----END CERTIFICATE-----`, v.X5c[0])

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("can't parse pem %q", pem)
		}
		certs[v.Kid] = publicKey
	}

	return certs, nil
}
