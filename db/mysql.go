package db

import (
	"database/sql"

	// Register MySQL
	_ "github.com/go-sql-driver/mysql"
)

// DB Conn
var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@/wiki?charset=utf8")
	DB = db
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
