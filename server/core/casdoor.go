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
		"c68223c4e2195dbe987f",
		"33ed90a877586dc435a233cce2930f8ac42623af",
		JwtPublicKey,
		"built-in",
		"usercenter")
	return nil
}
