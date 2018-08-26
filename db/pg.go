package db

import (
	"database/sql"
	"fmt"

	// Register PostgreSQL
	_ "github.com/lib/pq"
)

const (
	// DBUser 用户名
	DBUser = "root"
	// DBPassword 密码
	DBPassword = "123456"
	// DBName  库名
	DBName = "wiki"
)

// PostgreSQLDB Conn
var PostgreSQLDB *sql.DB

func init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)

	PostgreSQLDB = db
	checkPostgreSQLErr(err)
}

// Create 添加一条记录
func Create(title string) {
	var lastInsertID int
	err := PostgreSQLDB.QueryRow("INSERT INTO blog(title) VALUES($1) returning autokid;", title).Scan(&lastInsertID)
	checkPostgreSQLErr(err)
	fmt.Println("last inserted id =", lastInsertID)
}

func checkPostgreSQLErr(err error) {
	if err != nil {
		panic(err)
	}
}
