package selectPost

import (
	"net/http"
	"fmt"
	"time"
	"JWT/token/sign"
	"JWT/JsonResponse"
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


func Select(w http.ResponseWriter, r *http.Request) {

	postInfo := make(map[string]string)
	postInfo["name"] = r.Header.Get("name")
	fmt.Printf("post中的name=%v\n", r.Header.Get("name"))
	claims := &MyClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		name:      r.PostFormValue("name"),
	}

	fmt.Println(*claims)

	tokenString := sign.Sign(postInfo, "secret")

	response := Token{tokenString}
	JsonResponse.JsonResponse(response, w)

}
