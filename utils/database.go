package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = ""
	dbname = "travel"
)

var DB *sql.DB

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbport, dbuser, dbpass, dbname)

	fmt.Println(psqlInfo)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	rows, err := DB.Query("SELECT current_database();")

	for rows.Next() {
		var current_database string
		err = rows.Scan(&current_database)
		fmt.Println("DB Connected to:", current_database)
	}
}
