package LogOffResult

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"

	"JWT/token/VerifyToken"
	"JWT/JsonResponse"
	"JWT/checkErr"
	"database/sql"
)

type Response struct {
	Data string `json:"data"`
}

func LogOffResult(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Header.Get("Authorization"))
	ss := r.Header.Get("Authorization")
	var response interface{}
	claims, _ := VerifyToken.VerifyToken(ss, "secret")

	if claims1, ok := claims.(jwt.MapClaims); ok { //取出token中的数据
		fmt.Printf("username=%v\n", claims1["username"])
		username, ok1 := claims1["username"].(string)
		if username == "" {
			response = Response{"您还未登录"}
			JsonResponse.JsonResponse(response, w)
			return
		}
		if ok1 {
			ok3 := logOff(username)
			if !ok3 {
				response = Response{"注销失败"}
				JsonResponse.JsonResponse(response, w)
			} else {
				response = Response{"注销成功，您以后不能以此账户登录了"}
				JsonResponse.JsonResponse(response, w)
			}
		}
	}
}

func logOff(username string) bool {
	db, err := sql.Open("mysql", "root:20111412e@tcp(172.105.204.252:3306)/PXTest?charset=utf8")
	checkErr.CheckErr(err)
	defer db.Close()
	_, offErr := db.Exec("UPDATE UserInfo set state = 'false' where username='" + username + "'")
	if offErr != nil {
		fmt.Println("注销失败")
		return false
	}
	fmt.Println("注销成功")
	return true
}