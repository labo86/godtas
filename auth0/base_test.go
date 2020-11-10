package auth0

/*
 * las llaves fueron generadas en https://jwt.io/
 */
import (
	"fmt"
	"github.com/labo86/godtas/auth0test"
	"net/http/httptest"
	"testing"
)

func Test_ParseJsonKeys(t *testing.T) {
	config := Config(auth0test.LoadConfig())
	auth, err := config.Init()

	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
	}

	middleware := auth.MiddleWare()

	{
		r := httptest.NewRequest("GET", "/", nil)
		auth0test.SetToken(r)
		w := httptest.NewRecorder()

		if err := middleware.CheckJWT(w, r); err != nil {
			t.Errorf("fallo el jwt: %v", err)
		}

	}

	{
		token := `asfadfasdfasd`
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		w := httptest.NewRecorder()

		if err := middleware.CheckJWT(w, r); err == nil {
			t.Errorf("deberia fallar porque es un token random: %v", err)
		}
	}

	{
		token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImNlcnRfa2V5XzIifQ.eyJzdWIiOiJ0ZXN0fDEyMzQ1Njc4OTAiLCJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.jK4mtYxDjH8KGxaxk9W7uvpjWn2mGs_10L8HcH6pGCP8mSJOOzz2kR7JsdDBbo_LCE6xdPdclkPzHgp_xkGUZ71olp9rX8mPGSZkMeg-ihFU12JvCwn-JSIVreKH_UkA7giVKxdVfuMOrD30M4NiVRciebldw3pmZw5HxH67Hl1QJIllggnvOzHeta3lLKqtNmV08F-bKTz6U2cMh3OQU75f77XkA3biQ3E-F7sUhg7vOiT4W8KzfUY5SnCKNYAo3WpeJ2HUQXcmE_hE2I_bjNfl48Iifev4tscuQYd911ot7Hg_fwne9IfRPAkTc8Pv9kMUGG5fvIjEI0i6zX8v_w`
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		w := httptest.NewRecorder()

		if err := middleware.CheckJWT(w, r); err == nil {
			t.Errorf("deberia fallar porque tiene otra kid: %v", err)
		}
	}
}

func TestAuth0User(t *testing.T) {
	config := Config(auth0test.LoadConfig())
	auth, err := config.Init()

	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
	}

	middleware := auth.MiddleWare()

	{
		r := httptest.NewRequest("GET", "/", nil)
		auth0test.SetToken(r)
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
