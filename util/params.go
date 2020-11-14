package util

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labo86/godtas/service/auth0"
	"mime/multipart"
	"net/http"
	"strconv"
)

type Params struct {
	r     *http.Request
	Error error
}

func NewParams(r *http.Request) *Params {
	return &Params{
		r: r,
	}
}

func (p *Params) Route(name string) string {
	if !p.Ok() {
		return ""
	}

	vars := mux.Vars(p.r)
	value, ok := vars[name]
	if !ok {
		p.Error = fmt.Errorf("param %q not defined", name)
		return ""
	}
	return value
}

func (p *Params) User() string {
	if !p.Ok() {
		return ""
	}

	value, err := auth0.RequestUser(p.r)

	if err != nil {
		p.Error = err
	}
	return value
}

func (p *Params) FormValue(name string) string {
	if !p.Ok() {
		return ""
	}

	return p.r.FormValue(name)
}

func (p *Params) FormInt(name string) int {
	if !p.Ok() {
		return 0
	}

	value, err := strconv.Atoi(p.r.FormValue(name))
	if err != nil {
		p.Error = err
	}
	return value
}

func (p *Params) FormFile(name string) (multipart.File, *multipart.FileHeader) {
	if !p.Ok() {
		return nil, nil
	}

	value, headers, err := p.r.FormFile(name)
	if err != nil {
		p.Error = err
	}
	return value, headers
}

func (p *Params) Ok() bool {
	return p.Error == nil
}

func (p *Params) IsWrong(w http.ResponseWriter) bool {
	if !p.Ok() {
		http.Error(w, fmt.Sprintf("wrong params: %v", p.Error), http.StatusBadRequest)
		return true
	}
	return false
}
