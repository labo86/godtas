package util

import (
	"github.com/labo86/godtas/service/auth0"
	"net/http/httptest"
	"testing"
)

func TestParams_Ok(t *testing.T) {
	auth, err := auth0.NewTmp()

	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
	}

	middleware := auth.MiddleWare()

	{
		r := httptest.NewRequest("GET", "/", nil)
		auth0.SetTokenTest(r)
		w := httptest.NewRecorder()

		if err := middleware.CheckJWT(w, r); err != nil {
			t.Errorf("fallo el jwt: %v", err)
		}

		p := NewParams(r)
		user := p.User()

		expectedUser := "test|1234567890"
		if user != expectedUser {
			t.Errorf("usuario incorrecto %q deberia ser %q", user, expectedUser)
		}

		if p.IsWrong(w) {
			t.Errorf("should be ok")
		}

	}
}

func TestParamHelper_OK_NoOK(t *testing.T) {
	{
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		p := NewParams(r)
		user := p.User()

		expectedUser := ""
		if user != expectedUser {
			t.Errorf("el usuario debe ser vacio %q", user)
		}

		if !p.IsWrong(w) {
			t.Errorf("debe estar malo")
		}

	}
}
