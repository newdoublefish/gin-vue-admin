package core

import (
	_ "embed"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

type AuthResponse struct {
	Code string `json:"data"`
}

//go:embed token_jwt_key.pem
var JwtPublicKey string

var endPoint = "http://10.0.0.69:8000"

//TODO: read from config file

func Casdoor() error {
	auth.InitConfig(endPoint,
		"3dd9fd33b45821f80651",
		"a9ae00d7b70447922bb90841fcba794b33c2e239",
		JwtPublicKey,
		"built-in",
		"application_rr5d06")
	return nil
}
