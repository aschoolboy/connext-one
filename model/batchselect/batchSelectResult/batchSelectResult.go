package batchSelectResult

import (
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"JWT/token/VerifyToken"
	"JWT/_sql/batchQueryInfo"
	"JWT/JsonResponse"
)

type Response struct {
	Data string `json:"data"`
}

func BatchSelectResult(w http.ResponseWriter, r *http.Request) {

	var str string
	ss := r.Header.Get("Authorization")
	var response interface{}
	claims, _ := VerifyToken.VerifyToken(ss, "secret")

	fmt.Printf("解析后的token：%v\n", claims)
	if claims1, ok := claims.(jwt.MapClaims); ok { //取出token中的数据
		fmt.Printf("name=%v,num=%v\n", claims1["name"], claims1["num"])
		name, ok1 := claims1["name"].(string)
		num, ok2 := claims1["num"].(string)

		fmt.Printf("name:%T,num:%T\n", name, num)
		fmt.Println(ok1, ok2)
		fmt.Printf("name:%v,num:%v\n", name, num)

		if ok1 && ok2 {
			iNum, _ := strconv.Atoi(num)
			fmt.Printf("iNum=%v", iNum)
			users := batchQueryInfo.BatchQueryInfo(name, iNum)
			for i := 0; i < iNum; i++ {
				str = str + "{name=" + users[i].Name + ",sex=" + users[i].Sex + "}\n"
			}

			response = Response{str}
			JsonResponse.JsonResponse(response, w)
		}
	}
}