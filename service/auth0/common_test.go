package auth0

import (
	"testing"
)

func Test_ParseJsonKeys(t *testing.T) {
	auth, err := NewTmp()
	if err != nil {
		t.Errorf("el parseo debio ser exitoso :%v", err)
		return
	}

	{
		if _, err := auth.CheckJWT(TokenTest); err != nil {
			t.Errorf("fallo el jwt: %v", err)
			return
		}
	}

	{
		token := `asfadfasdfasd`

		if _, err := auth.CheckJWT(token); err == nil {
			t.Errorf("deberia fallar porque es un token random: %v", err)
		}
	}

	{
		token := `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImNlcnRfa2V5XzIifQ.eyJzdWIiOiJ0ZXN0fDEyMzQ1Njc4OTAiLCJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.jK4mtYxDjH8KGxaxk9W7uvpjWn2mGs_10L8HcH6pGCP8mSJOOzz2kR7JsdDBbo_LCE6xdPdclkPzHgp_xkGUZ71olp9rX8mPGSZkMeg-ihFU12JvCwn-JSIVreKH_UkA7giVKxdVfuMOrD30M4NiVRciebldw3pmZw5HxH67Hl1QJIllggnvOzHeta3lLKqtNmV08F-bKTz6U2cMh3OQU75f77XkA3biQ3E-F7sUhg7vOiT4W8KzfUY5SnCKNYAo3WpeJ2HUQXcmE_hE2I_bjNfl48Iifev4tscuQYd911ot7Hg_fwne9IfRPAkTc8Pv9kMUGG5fvIjEI0i6zX8v_w`

		if _, err := auth.CheckJWT(token); err == nil {
			t.Errorf("deberia fallar porque tiene otra kid: %v", err)
		}
	}
}

func TestClaimValue(t *testing.T) {

	_, err := ClaimValue(nil, "some")
	if err == nil {
		t.Errorf("deberia fallar por claim")
		return
	}

}
