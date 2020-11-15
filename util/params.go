package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type Params struct {
	r      *http.Request
	Errors []error
}

func NewParams(r *http.Request) *Params {
	return &Params{
		r: r,
	}
}

func (p *Params) Route(name string) string {
	vars := mux.Vars(p.r)
	value, ok := vars[name]
	if !ok {
		p.Errors = append(p.Errors, fmt.Errorf("route param %q not defined", name))
		return ""
	}
	return value
}

func (p *Params) AuthorizationToken() string {
	authHeader := p.r.Header.Get("Authorization")

	if authHeader == "" {
		p.Errors = append(p.Errors, errors.New("empty authorization bearer"))
		return ""
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		p.Errors = append(p.Errors, fmt.Errorf("authorization header format must be Bearer {token} : %q", authHeader))
		return ""
	}

	token := authHeaderParts[1]

	if token == "" {
		p.Errors = append(p.Errors, fmt.Errorf("token is empty : %q", authHeader))
	}

	return token
}

func (p *Params) FormValue(name string) string {
	return p.r.FormValue(name)
}

func (p *Params) FormInt(name string) int {
	value, err := strconv.Atoi(p.r.FormValue(name))
	if err != nil {
		p.Errors = append(p.Errors, fmt.Errorf("form value %q is not int : %v", name, err))
	}
	return value
}

func (p *Params) FormFile(name string) (multipart.File, *multipart.FileHeader) {
	value, headers, err := p.r.FormFile(name)
	if err != nil {
		p.Errors = append(p.Errors, fmt.Errorf("can't obtain form file value %q: %v", name, err))
	}
	return value, headers
}

func (p *Params) Ok() bool {
	return len(p.Errors) == 0
}

func (p *Params) JSON(value interface{}) {
	contentType := p.r.Header.Get("Content-type")
	if want := "application/json"; !strings.HasPrefix(contentType, want) {
		p.Errors = append(p.Errors, fmt.Errorf("content type not a json : actual %q", contentType))
		return
	}

	if err := json.NewDecoder(p.r.Body).Decode(value); err != nil {
		p.Errors = append(p.Errors, fmt.Errorf("can't decode body as json: %v", err))
		return
	}
}

func (p *Params) IsWrong(w http.ResponseWriter) bool {
	if !p.Ok() {
		http.Error(w, fmt.Sprintf("wrong params: %v", p.Errors), http.StatusBadRequest)
		return true
	}
	return false
}
