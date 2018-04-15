package batchQueryInfo

import (
	"database/sql"
	"fmt"
	"JWT/checkErr"
)

type User struct {
	UserName string
	PassWord string
	Name     string
	Sex      string
	state    string
}

var Users []User

func BatchQueryInfo(name string, num int) []User {
	db, err := sql.Open("mysql", "root:20111412e@tcp(172.105.204.252:3306)/PXTest?charset=utf8")
	checkErr.CheckErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM UserInfo where name='" + name + "'")
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

		}
		fmt.Println("-----------------------------------")
		Users = append(Users, User{userInfo["username"], "", userInfo["name"], userInfo["sex"], userInfo["state"]})
		// return User{userInfo["name"], userInfo["username"], userInfo["sex"]}

		fmt.Println("-----------------------------------")
	}
	return Users
}
