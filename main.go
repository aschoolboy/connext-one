package main

import (
	"log"
	"net/http"
	"JWT/model/login/loginPost"
	"JWT/model/_select/selectPost"
	"JWT/model/logout/logOutPost"
	"JWT/model/batchselect/batchSelectPost"
	"JWT/model/logoff/logoffPost"
	"JWT/model/login/loginResult"
	"JWT/model/_select/SelectResult"
	"JWT/model/logout/LogOutResult"
	"JWT/model/batchselect/batchSelectResult"
	"JWT/model/logoff/LogOffResult"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	StartServer()
}

func StartServer() {

	http.HandleFunc("/login", loginPost.LoginPost)
	http.HandleFunc("/select", selectPost.Select)
	http.HandleFunc("/batchSelect", batchSelectPost.BatchSelect)
	http.HandleFunc("/logOff", logoffPost.LogOff)
	http.HandleFunc("/logOut", logOutPost.LogOut)

	http.HandleFunc("/loginResult", loginResult.LoginResult)
	http.HandleFunc("/selectResult", SelectResult.SelectResult)
	http.HandleFunc("/batchSelectResult", batchSelectResult.BatchSelectResult)
	http.HandleFunc("/logOffResult", LogOffResult.LogOffResult)
	http.HandleFunc("/logOutResult", LogOutResult.LogOutResult)

	log.Println("Now listening...")
	http.ListenAndServe(":8080", nil)

}

