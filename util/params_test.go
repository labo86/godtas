package util

import (
	"github.com/labo86/godtas/service/auth0"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParams_Ok(t *testing.T) {

	{
		r := httptest.NewRequest("GET", "/", nil)
		auth0.SetTokenTest(r)

		p := NewParams(r)
		token := p.AuthorizationToken()

		if got, want := token, auth0.TokenTest; got != want {
			t.Errorf("token incorrecto %q deberia ser %q", got, want)
			return
		}

		if err := p.Error(); err != nil {
			t.Errorf("should be ok: %v", err)
			return
		}

	}
}

func TestParamHelper_OK_NoOK(t *testing.T) {
	{
		r := httptest.NewRequest("GET", "/", nil)

		p := NewParams(r)
		token := p.AuthorizationToken()

		if got, want := token, ""; got != want {
			t.Errorf("token incorrecto %q deberia ser %q", got, want)
			return
		}

		if err := p.Error(); err == nil {
			t.Errorf("debe estar malo")
			return
		}

	}
}

func TestParamHelper_JSON(t *testing.T) {
	{
		body := strings.NewReader(`{"name":"hola"}`)

		r := httptest.NewRequest("POST", "/", body)
		r.Header.Set("Content-Type", "application/json")

		var value map[string]string
		p := NewParams(r)
		p.JSON(&value)

		if err := p.Error(); err != nil {
			t.Errorf("debe ser un json valido : %v", p.Error())
			return
		}

		key := "name"
		sub, ok := value[key]
		if !ok {
			t.Errorf("key %q should exists : %v", key, value)
			return
		}

		if got, want := sub, "hola"; got != want {
			t.Errorf("value of %q got %q, want %q", key, got, want)
			return
		}

	}
}

func TestParamHelper_JSON2(t *testing.T) {
	{
		body := strings.NewReader(`{"name":"hola"}`)

		r := httptest.NewRequest("POST", "/", body)
		r.Header.Set("Content-Type", "application/json; charset=utf-8")

		var value map[string]string
		p := NewParams(r)
		p.JSON(&value)

		if err := p.Error(); err != nil {
			t.Errorf("debe ser un json valido : %v", p.Error())
			return
		}

		key := "name"
		sub, ok := value[key]
		if !ok {
			t.Errorf("key %q should exists : %v", key, value)
			return
		}

		if got, want := sub, "hola"; got != want {
			t.Errorf("value of %q got %q, want %q", key, got, want)
			return
		}

	}
}

func TestParamHelper_JSONInvalidFormat(t *testing.T) {
	{
		body := strings.NewReader(`{{`)

		r := httptest.NewRequest("POST", "/", body)
		r.Header.Set("Content-Type", "application/json; charset=utf-8")

		var value map[string]string
		p := NewParams(r)
		p.JSON(&value)

		if err := p.Error(); err == nil {
			t.Errorf("no debe ser un json valido")
			return
		}

	}
}

func TestParamHelper_JSONInvalidContentType(t *testing.T) {
	{
		body := strings.NewReader(`{"name":"hola"}`)

		r := httptest.NewRequest("POST", "/", body)
		r.Header.Set("Content-Type", "adfasdf")

		var value map[string]string
		p := NewParams(r)
		p.JSON(&value)

		if err := p.Error(); err == nil {
			t.Errorf("no debe ser un json valido")
			return
		}

	}
}

func TestParams_Auth0ClaimOk(t *testing.T) {

	{
		a, _ := auth0.NewTmp()

		r := httptest.NewRequest("GET", "/", nil)
		auth0.SetTokenTest(r)

		p := NewParams(r)
		sub := p.Auth0Claim(a, "sub")

		if got, want := sub, "test|1234567890"; got != want {
			t.Errorf("sub incorrecto %q deberia ser %q", got, want)
			return
		}

		if err := p.Error(); err != nil {
			t.Errorf("should be ok: %v", err)
			return
		}

		sub = p.Auth0Claim(a, "sub")

		if got, want := sub, "test|1234567890"; got != want {
			t.Errorf("sub incorrecto %q deberia ser %q", got, want)
			return
		}

		if err := p.Error(); err != nil {
			t.Errorf("should be ok: %v", err)
			return
		}

		sub = p.Auth0Claim(a, "not_existant")

		if err := p.Error(); err == nil {
			t.Errorf("deberia fallar")
			return
		}
	}
}

func TestParams_OkServiceError(t *testing.T) {

	{
		r := httptest.NewRequest("GET", "/", nil)
		auth0.SetTokenTest(r)

		p := NewParams(r)
		token := p.AuthorizationToken()

		if got, want := token, auth0.TokenTest; got != want {
			t.Errorf("token incorrecto %q deberia ser %q", got, want)
			return
		}

		if err := p.ServiceError(); err != nil {
			t.Errorf("should be ok: %v", err)
			return
		}

	}
}
