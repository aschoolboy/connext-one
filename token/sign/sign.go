package sign

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"

)

func Sign(info map[string]string, secret string) string { //生成token,算法hs256

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["name"] = info["name"]
	claims["secret"] = secret //需私钥加密
	claims["username"] = info["username"]
	claims["password"] = info["password"]
	claims["state"] = info["state"]
	claims["num"] = info["num"]
	token.Claims = claims
	s, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	return s
}
