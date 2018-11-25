package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
func AddUser(firstname string, lastname string, userID string, username string) {

	database, error := sql.Open("sqlite3", "./database.db?cache=shared&mode=rwc")
	if error != nil {
		fmt.Println("Ошибка")
		fmt.Println(error)
	}

	statement, error := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, userID TEXT, username TEXT)")
	if error != nil {
		fmt.Println("Ошибка")
		fmt.Println(error)
	}
	statement.Exec()

    statement, error = database.Prepare("INSERT INTO users (firstname, lastname, userID, username) VALUES (?,?,?,?)")
	if error != nil {
		fmt.Println("Ошибка")
		fmt.Println(error)
	}
	statement.Exec(firstname, lastname, userID, username)

}




/*func ShowUser(IdOfUser int)  {

	database, error := sql.Open("sqlite3", "./database.db?cache=shared&mode=rwc")
	if error != nil {
		fmt.Println("Ошибка")
		fmt.Println(error)
	}
	 __,error := database.Query("SELECT userID,firstname,lastname FROM users WHERE userID = IdOfUser")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
		fmt.Println(__)
	}
}*/





