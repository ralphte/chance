package sql

import (
        _"github.com/go-sql-driver/mysql"
        "database/sql"
)

func GetUserPassword(username string) string {

        db, err := sql.Open("mysql", "chance_app:Cu46qNDytt2T@tcp(localhost:3306)/chance?charset=utf8")
        checkErr(err)
        defer db.Close()

        rows, err := db.Query("SELECT password FROM users WHERE username = ? LIMIT 1", username)
        checkErr(err)

        var password string

        for rows.Next() {
                err := rows.Scan(&password)
                checkErr(err)
        }
        return password
}

func GetUserData(username string) (string, string, string) {

        db, err := sql.Open("mysql", "chance_app:Cu46qNDytt2T@tcp(localhost:3306)/chance?charset=utf8")
        checkErr(err)
        defer db.Close()

        rows, err := db.Query("SELECT firstname, lastname, email FROM users WHERE username = ? LIMIT 1", username)
        checkErr(err)

        var firstname string
        var lastname string
        var email string

        for rows.Next() {
                err := rows.Scan(&firstname, &lastname, &email)
                checkErr(err)
        }
        return email, firstname, lastname
}

/*func UpdateUserData(username string, firstname string, lastname string, email string) (string) {

	db, err := sql.Open("mysql", "chance_app:Cu46qNDytt2T@tcp(localhost:3306)/chance?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Prepare("update users set firstname, lastname, email set username=?", username)
	checkErr(err)

	var firstname string
	var lastname string
	var email string

	for rows.Next() {
		err := rows.Scan(&firstname, &lastname, &email)
		checkErr(err)
	}
	return email, firstname, lastname
}
*/


func checkErr(err error) {
        if err != nil {
                panic(err)
        }
}