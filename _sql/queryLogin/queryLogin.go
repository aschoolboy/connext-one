package queryLogin

import (
	"fmt"

	"JWT/checkErr"
	"database/sql"
)

type User struct {
	UserName string
	PassWord string
	Name     string
	Sex      string
	State    string
}

func QueryLogin(username, password string) (bool, User) {
	db, err := sql.Open("mysql", "root:20111412e@tcp(172.105.204.252:3306)/PXTest?charset=utf8")
	checkErr.CheckErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM UserInfo")
	checkErr.CheckErr(err)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	userInfo := make(map[string]string)

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			userInfo[columns[i]] = value
			// userInfos=append(userInfos,userInfo)

		}
		fmt.Println("-----------------------------------")
		if userInfo["username"] == username && userInfo["password"] == password && userInfo["state"] == "true" {
			return true, User{userInfo["username"], "", userInfo["name"], userInfo["sex"], userInfo["state"]}
		}

		fmt.Println("-----------------------------------")
	}
	return false, User{"", "", "", "", ""}
}