package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/labo86/godtas/service/auth0"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type Params struct {
	Req    *http.Request
	Errors ErrorList
	Values map[string]interface{}
}

func NewParams(r *http.Request) *Params {
	return &Params{
		Req:    r,
		Values: make(map[string]interface{}),
	}
}

func (p *Params) Route(name string) string {
	vars := mux.Vars(p.Req)
	value, ok := vars[name]
	if !ok {
		p.Errors.Append(fmt.Errorf("route param %q not defined", name))
		return ""
	}
	return value
}

func (p *Params) AuthorizationToken() string {
	authHeader := p.Req.Header.Get("Authorization")

	if authHeader == "" {
		p.Errors.Append(errors.New("empty authorization bearer"))
		return ""
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		p.Errors.Append(fmt.Errorf("authorization header format must be Bearer {token} : %q", authHeader))
		return ""
	}

	token := authHeaderParts[1]

	if token == "" {
		p.Errors.Append(fmt.Errorf("token is empty : %q", authHeader))
	}

	return token
}

func (p *Params) FormValue(name string) string {
	return p.Req.FormValue(name)
}

func (p *Params) FormInt(name string) int {
	value, err := strconv.Atoi(p.Req.FormValue(name))
	if err != nil {
		p.Errors.Append(fmt.Errorf("form value %q is not int : %v", name, err))
	}
	return value
}

func (p *Params) FormFile(name string) (multipart.File, *multipart.FileHeader) {
	value, headers, err := p.Req.FormFile(name)
	if err != nil {
		p.Errors.Append(fmt.Errorf("can't obtain form file value %q: %v", name, err))
	}
	return value, headers
}

func (p *Params) Error() error {
	return p.Errors.FirstError()
}

func (p *Params) ServiceError() error {
	if err := p.Errors.FirstError(); err != nil {
		return NewServiceError("", http.StatusBadRequest, p.Errors.FirstError())
	}
	return nil
}

func (p *Params) JSON(value interface{}) {
	contentType := p.Req.Header.Get("Content-type")
	if want := "application/json"; !strings.HasPrefix(contentType, want) {
		p.Errors.Append(fmt.Errorf("content type not a json : actual %q", contentType))
		return
	}

	if err := json.NewDecoder(p.Req.Body).Decode(value); err != nil {
		p.Errors.Append(fmt.Errorf("can't decode body as json: %v", err))
		return
	}
}

func (p *Params) Auth0AuthorizationToken(a auth0.Auth0) *jwt.Token {
	value, ok := p.Values["auth0_authorization_token"].(*jwt.Token)
	if ok {
		return value
	}

	token := p.AuthorizationToken()

	parsedToken, err := a.CheckJWT(token)
	if err != nil {
		p.Errors.Append(fmt.Errorf("authorization token error : %v", err))
		return nil
	}

	p.Values["auth0_authorization_token"] = parsedToken
	return parsedToken

}

func (p *Params) Auth0Claim(a auth0.Auth0, name string) string {

	token := p.Auth0AuthorizationToken(a)

	value, err := auth0.ClaimValue(token, name)
	if err != nil {
		p.Errors.Append(fmt.Errorf("claim %q : %v", name, err))
		return ""
	}

	return value
}
