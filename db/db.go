package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = "goapi"
)

func OpenConnection() *sql.DB {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable",
		DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	// defer db.Close()
	return db
}

func GetAllUsers(db *sql.DB) []byte {
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	data := []byte(`[`)
	isNotLast := rows.Next()
	for isNotLast {
		var name string
		var email string
		var password string
		err = rows.Scan(&name, &email, &password)
		checkErr(err)
		fmt.Println("name | email | password")
		fmt.Printf("%3v | %8v | %6v\n", name, email, password)
		row := fmt.Sprintf(`{"name": "%s", "email": "%s", "password": "%s"}`, name, email, password)
		isNotLast = rows.Next()
		data = append(data, row...)
		if isNotLast {
			data = append(data, ", "...)
		}
	}
	data = append(data, "]"...)
	defer db.Close()
	return data
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
