package utils

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() *sql.DB {
	usuario := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", usuario, pass, host, dbname))
	if err != nil {
		panic(err)
	}

	return db
}
