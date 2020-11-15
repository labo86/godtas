package auth0

import (
	"testing"
)

func TestTmp(t *testing.T) {
	auth, err := NewTmp()

	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
	}

	token, err := auth.CheckJWT(TokenTest)
	if err != nil {
		t.Errorf("fallo el jwt: %v", err)
		return
	}

	value, err := ClaimValue(token, "sub")

	if err != nil {
		t.Errorf("no se obtuvo el user: %v", err)
	}

	if got, want := value, "test|1234567890"; got != want {
		t.Errorf("usuario incorrecto %q deberia ser %q", got, want)
	}

}
