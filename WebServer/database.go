package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Printf("Creat database")
	database, _ :=
		sql.Open("sqlite3", "./users.db")
	fmt.Println("---> done")
	fmt.Printf("Create table user")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password INTEGER,timecreat DATETIME CURRENT_TIMESTAMP)")
	statement.Exec()
	fmt.Println("---> done")
	fmt.Printf("Insert data into table user")
	statement, err :=
		database.Prepare("INSERT INTO users (username, password,timecreat) VALUES (?, ?,?)")
	if err != nil {
		fmt.Printf("Error : ")
		fmt.Println(err)
	}
	statement.Exec("Haimd", "123456", "datetime()")
	fmt.Println("---> done")
	fmt.Printf("Query data into table user")
	rows, _ :=
		database.Query("SELECT id, username, password, timecreat FROM users")
	fmt.Println("---> done")
	var id int
	var user string
	var pass string

	for rows.Next() {
		rows.Scan(&id, &user, &pass)
		fmt.Println(strconv.Itoa(id) + ": " + user + " " + pass)
	}

}
