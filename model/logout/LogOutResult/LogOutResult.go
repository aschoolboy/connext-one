package LogOutResult

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"

	"JWT/token/VerifyToken"
	"JWT/JsonResponse"
)

type Response struct {
	Data string `json:"data"`
}

type User struct {
	UserName string
	PassWord string
	Name     string
	Sex      string
	state    string
}

var UserCache = User{"", "", "", "", ""}

func LogOutResult(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Header.Get("Authorization"))
	ss := r.Header.Get("Authorization")
	var response interface{}
	claims, _ := VerifyToken.VerifyToken(ss, "secret")

	if claims1, ok := claims.(jwt.MapClaims); ok { //取出token中的数据
		fmt.Printf("username=%v\n", claims1["username"])
		username, ok1 := claims1["username"].(string)
		if username != "" && ok1 == true {
			UserCache = User{"", "", "", "", ""}
			response = Response{"您已成功登出"}
			JsonResponse.JsonResponse(response, w)
			return
		}

	}
	response = Response{"请先登录"}
	JsonResponse.JsonResponse(response, w)
}