package SelectResult

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"JWT/token/VerifyToken"
	"JWT/_sql/queryInfo"
	"JWT/JsonResponse"
)

type Response struct {
	Data string `json:"data"`
}

func SelectResult(w http.ResponseWriter, r *http.Request) {

	ss := r.Header.Get("Authorization")
	var response interface{}
	claims, _ := VerifyToken.VerifyToken(ss, "secret")

	if claims1, ok := claims.(jwt.MapClaims); ok { //取出token中的数据
		fmt.Printf("name=%v\n", claims1["name"])
		name, ok1 := claims1["name"].(string)

		if ok1 {
			user := queryInfo.QueryInfo(name)
			str := "name=" + user.Name + ",sex=" + user.Sex
			response = Response{str}
			JsonResponse.JsonResponse(response, w)
		}
	}
}