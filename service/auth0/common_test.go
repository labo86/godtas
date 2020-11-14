package auth0

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func Test_ParseJsonKeys(t *testing.T) {
	auth, err := NewTmp()
	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
		return
	}

	middleware := auth.MiddleWare()

	{
		r := httptest.NewRequest("GET", "/", nil)
		SetTokenTest(r)
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
