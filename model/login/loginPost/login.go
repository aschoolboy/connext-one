package loginPost

import (
	"net/http"
	"JWT/token/sign"
	"JWT/JsonResponse"
)

type Token struct {
	Token string `json:"token"`
}

func LoginPost(w http.ResponseWriter, r *http.Request){
	postInfo := make(map[string]string)

	postInfo["username"] = r.PostFormValue("username")
	postInfo["password"] = r.PostFormValue("password")

	tokenString := sign.Sign(postInfo, "secret")
	response := Token{tokenString}
	JsonResponse.JsonResponse(response, w)
}