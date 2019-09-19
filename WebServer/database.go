package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Printf("Creat database")
	database, _ :=
		sql.Open("sqlite3", "./users.db")
	fmt.Println("---> done")
	fmt.Printf("Create table user")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(64) NULL, password VARCHAR(64) NULL, timecreate DATE NULL)")
	statement.Exec()
	fmt.Println("---> done")
	fmt.Printf("Insert data into table user")
	statement, err :=
		database.Prepare("INSERT INTO users (username, password,timecreate) VALUES (?, ?,?)")
	if err != nil {
		fmt.Printf("Error : ")
		fmt.Println(err)
	}
	statement.Exec("Haimd", "123456", "2012-12-09")
	fmt.Println("---> done")
	fmt.Printf("Query data into table users")
	rows, err :=
		database.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---> done")

		var uid int
		var username string
		var pass string
		var created time.Time

		for rows.Next() {
			err = rows.Scan(&uid, &username, &pass, &created)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print(uid)
				fmt.Print("\t")
				fmt.Print(username)
				fmt.Print("\t")
				fmt.Print(pass)
				fmt.Print("\t")
				fmt.Println(created)
			}

		}

		rows.Close() //good habit to close

	}

}
