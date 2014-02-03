package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main123() {

	db, err := sql.Open("mysql", "root:Password1@tcp(127.0.0.1:3307)/world")

	checkError := func() {
		if err != nil {
			//fmt.Println(err.Error())
			panic(err.Error())
		}
	}

	checkError()
	r, err := db.Query("select Code from country")
	checkError()
	for r.Next() {
		var code string
		r.Scan(&code)
		fmt.Println(code)
	}
	//fmt.Println(db)
	db.Close()
}
