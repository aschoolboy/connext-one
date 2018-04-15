package loginResult

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"JWT/token/VerifyToken"
	"JWT/_sql/queryLogin"
	"JWT/JsonResponse"
)

var UserCache = User{"", "", "", "", ""}

type User struct {
	UserName string
	PassWord string
	Name     string
	Sex      string
	state    string
}

type Response struct {
	Data string `json:"data"`
}

func LoginResult(w http.ResponseWriter, r *http.Request) {
	ss := r.Header.Get("Authorization")
	var response interface{}
	claims, _ := VerifyToken.VerifyToken(ss, "secret")

	if claims1, ok := claims.(jwt.MapClaims); ok { //取出token中的数据
		fmt.Printf("username=%v,password=%v\n", claims1["username"], claims1["password"])
		username, ok1 := claims1["username"].(string)
		password, ok2 := claims1["password"].(string)
		if ok1 && ok2 {
			ok3, user := queryLogin.QueryLogin(username, password)
			if !ok3 {
				response = Response{"登陆失败"}
				JsonResponse.JsonResponse(response, w)
			} else {
				UserCache = User{user.UserName, "", user.Name, user.Sex, user.State}
				response = Response{"登陆成功,欢迎您" + user.Name}
				JsonResponse.JsonResponse(response, w)
			}
		}
	}
}