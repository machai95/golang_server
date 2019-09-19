package main

import (
    "database/sql"
    "fmt"
    "strconv"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := 
		sql.Open("sqlite3", "./users.db")
	statement, _ := 
		database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password NUMBERIC,timecreat DATETIME CURRENT_TIMESTAMP)")
    statement.Exec()
	statement, _ = 
		database.Prepare("INSERT INTO users (username, password) VALUES (?, ?,?,?)")
     ("Haimd", "123456","datetime()")
	rows, _ := 
		database.Query("SELECT id, username, password, timecreat FROM users")
    var id int
    var user string
    var pass string

    for rows.Next() {
        rows.Scan(&id, &user, &pass)
        fmt.Println(strconv.Itoa(id) + ": " + user + " " + pass)
    }

}
