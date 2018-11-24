package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
func AddUser(firstname string, lastname string, userID string, username string) {

	database, error := sql.Open("sqlite3", "./database.db?cache=shared&mode=rwc")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}
	// database.Prepare("DROP TABLE  IF exists users")
	statement, error := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, userID TEXT, username TEXT)")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}
	statement.Exec()
    /*statement, error =database.Prepare("ALTER TABLE users ADD COLUMN firstname")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}
	statement.Exec()*/

    statement, error = database.Prepare("INSERT INTO users (firstname, lastname, userID, username) VALUES (?,?,?,?)")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}
	statement.Exec(firstname, lastname, userID, username)

	/*rows, _ := database.Query("SELECT id, firstname, lastname FROM users")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}*/
}

/*func ShowUser(IdOfUser int)  {

	database, error := sql.Open("sqlite3", "./database.db")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}

	statement,error := database.Query("SELECT  FROM userID WHERE userID=IdOfUser")
	if error != nil {
		fmt.Println("тут ошибка")
		fmt.Println(error)
	}
	statement.Exec()




 }*/
