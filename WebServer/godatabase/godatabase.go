package godatabase

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var uid int
var username string
var pass string
var created time.Time

func CreateDB() {
	fmt.Printf("Connect database")
	database, _ :=
		sql.Open("sqlite3", "./users.db")
	fmt.Println("---> done")
	fmt.Printf("Create table user if not exits")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(64) NULL, password VARCHAR(64) NULL, timecreate DATE NULL)")
	statement.Exec()
	fmt.Println("---> done")
	database.Close()
	//---------------- INSERT ----------------------
	// fmt.Printf("Insert data into table user")
	// statement, err :=
	// 	database.Prepare("INSERT INTO users (username, password,timecreate) VALUES (?, ?,?)")
	// if err != nil {
	// 	fmt.Printf("Error : ")
	// 	fmt.Println(err)
	// }
	// statement.Exec("Haimd", "123456", "2012-12-09")
	// fmt.Println("---> done")
}
func InsertDB(userclient string, passclient string) error {
	database, _ :=
		sql.Open("sqlite3", "./users.db")
	//---------------- INSERT ----------------------
	fmt.Printf("Insert data into table user")
	statement, err :=
		database.Prepare("INSERT INTO users (username, password,timecreate) VALUES (?, ?,?)")
	if err != nil {
		fmt.Printf("Error : ")
		fmt.Println(err)
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	statement.Exec(userclient, passclient, t)
	fmt.Println("---> done")
	return err
}
func CheckUser(userclient string, passclient string) bool {
	var stt = false
	database, _ :=
		sql.Open("sqlite3", "./users.db")
	fmt.Println("---> done")
	//----------------- QUERY ---------------------
	fmt.Printf("Query data into table users")
	rows, err :=
		database.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---> done")
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

				if username == userclient && pass == passclient {
					stt = true
					break
				}

			}

		}

		rows.Close() //good habit to close

	}
	database.Close()
	return stt
}
