package auth0

import (
	"net/http/httptest"
	"testing"
)

func TestTmp(t *testing.T) {
	auth, err := NewTmp()

	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
	}

	middleware := auth.MiddleWare()

	{
		r := httptest.NewRequest("GET", "/", nil)
		SetTokenTest(r)
		w := httptest.NewRecorder()

		if err := middleware.CheckJWT(w, r); err != nil {
			t.Errorf("fallo el jwt: %v", err)
		}

		user, err := RequestUser(r)
		if err != nil {
			t.Errorf("no se obtuvo el user: %v", err)
		}

		expectedUser := "test|1234567890"
		if user != expectedUser {
			t.Errorf("usuario incorrecto %q deberia ser %q", user, expectedUser)
		}
	}

}
