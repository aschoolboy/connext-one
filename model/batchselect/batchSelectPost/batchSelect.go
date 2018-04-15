package batchSelectPost

import (
	"net/http"
	"fmt"
	"time"

	"JWT/JsonResponse"
	"JWT/token/sign"
)

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

func BatchSelect(w http.ResponseWriter, r *http.Request) {

	postInfo := make(map[string]string)
	postInfo["name"] = r.Header.Get("name")
	postInfo["num"] = r.Header.Get("num")
	fmt.Printf("post中的name=%v,num=%v\n", r.Header.Get("name"), r.Header.Get("num"))
	claims := &MyClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		name:      r.PostFormValue("name"),
		num:       r.PostFormValue("num"),
	}

	fmt.Println(*claims)

	tokenString := sign.Sign(postInfo, "secret")

	response := Token{tokenString}
	JsonResponse.JsonResponse(response, w)

}