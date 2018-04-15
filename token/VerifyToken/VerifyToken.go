package VerifyToken

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"

	"JWT/getValidationKey"
)

func VerifyToken(input string, secret string) (jwt.Claims, string) {
	token, err := jwt.Parse(input, getValidationKey.GetValidationKey)
	if err != nil {
		return nil, err.Error()
	}
	if jwt.SigningMethodHS256.Alg() != token.Header["alg"] {
		return nil, "header err!"
	}

	fmt.Printf("verify pass!!!\n")
	return token.Claims, ""
}