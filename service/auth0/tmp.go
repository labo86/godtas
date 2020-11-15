package auth0

import (
	"fmt"
	"net/http"
)

type Tmp struct {
	Local
}

const TokenTest = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImNlcnRfa2V5In0.eyJzdWIiOiJ0ZXN0fDEyMzQ1Njc4OTAiLCJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwiaWF0IjoxNTE2MjM5MDIyfQ.NjeVg_iu6tI7ku2dfjzd1rg-taTrb2yZBjPwvZYspIKFdoA5EW63b30M2eYj9ydg1XOTh_ZstKThWK84XcjTihcw8hCecMVFmgjpH3gjBic3QHaS1u1TPKhNaT78Jf_qxMahriOFikxsojAYjuMR9-dMccYRMb_cMMZfa54Mu3kcQRwW842m3jBK0gLawoxo0FYUoXdYShjJOnlUTjdIIFCE-mqayxC4QMBwlBUmZfxjIDoPSYNfzKW5rANFFXn4wwKlWDkzDhrokKutWpw5ZkrHdsuxGRYFMdrv4AFNcNODLBmRRuDVSsdJqdYHQLT1k5x3h8-T3w8djulH1XZD_Q`

func NewTmp() (Auth0, error) {
	c := Config{
		Type: "local",
		Kid:  "cert_key",
		X5c: `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv
vkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc
aT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy
tvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0
e+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb
V6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9
MwIDAQAB`,
	}

	return c.Open()
}

func SetTokenTest(r *http.Request) {

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TokenTest))
}
