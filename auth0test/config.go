package auth0test

type Config struct {
	Type     string
	Audience string
	Issuer   string
	Url      string
	Kid      string
	X5c      string
}

func LoadConfig() Config {
	var config Config
	config.Type = "local"
	config.Issuer = "https://labo86.us.auth0.com/"
	config.Audience = "http://publicat.alpha.labo86.cl/api"
	config.Kid = "cert_key"
	config.X5c =
		`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv
vkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc
aT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy
tvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0
e+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb
V6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9
MwIDAQAB`

	return config
}
