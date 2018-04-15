package logoffPost

import (
	"net/http"
	"time"
	"fmt"

	"JWT/JsonResponse"

	"JWT/token/sign"
)

type User struct {
	UserName string
	PassWord string
	Name     string
	Sex      string
	state    string
}

var UserCache = User{"", "", "", "", ""}


type Token struct {
	Token string `json:"token"`
}

type MyClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	name      string
	username  string
	password  string
	sex       string
	num       string
}

func LogOff(w http.ResponseWriter, r *http.Request) {

	postInfo := make(map[string]string)
	postInfo["username"] = UserCache.UserName
	claims := &MyClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		username:  postInfo["username"],
	}

	fmt.Println(*claims)

	tokenString := sign.Sign(postInfo, "secret")

	response := Token{tokenString}
	JsonResponse.JsonResponse(response, w)

}