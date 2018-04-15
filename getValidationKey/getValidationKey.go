package getValidationKey

import "github.com/dgrijalva/jwt-go"

func GetValidationKey(*jwt.Token) (interface{}, error) {
	return []byte("secret"), nil
}