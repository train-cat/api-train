package security

import (
	"fmt"
	"log"

	"aahframework.org/aah.v0"
	"github.com/dgrijalva/jwt-go"
)

// JWT config values
var (
	JWTSigningMethod string
	JWTSigningKey    []byte
)

func init() {
	aah.OnStart(loadJWTConfig)
}

// CreateJWTToken return token based on config value.
// `security.auth_schemes.jwt_auth.signing_method` on security.conf
func CreateJWTToken() *jwt.Token {
	switch JWTSigningMethod {
	case "HS256":
		return jwt.New(jwt.SigningMethodHS256)
	case "HS384":
		return jwt.New(jwt.SigningMethodHS384)
	case "HS512":
		return jwt.New(jwt.SigningMethodHS512)
	}
	return nil
}

// jwtKeyFunc method to check the signing method
func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != JWTSigningMethod {
		return nil, fmt.Errorf("jwt: unexpected signing method '%v'", t.Header["alg"])
	}
	return JWTSigningKey, nil
}

// loads JWT setting into variable on server startup.
func loadJWTConfig(_ *aah.Event) {
	JWTSigningMethod = aah.AppConfig().StringDefault("security.auth_schemes.jwt_auth.signing_method", "HS256")

	if signingKey, found := aah.AppConfig().String("security.auth_schemes.jwt_auth.signing_key"); found {
		JWTSigningKey = []byte(signingKey)
	} else {
		log.Fatal("security.auth_schemes.jwt_auth.signing_key is not provided")
	}
}
