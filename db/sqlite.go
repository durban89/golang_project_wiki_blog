package db

import (
	"database/sql"
	"fmt"

	// Register SQLite
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteDB Conn
var SQLiteDB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "../data/data.db")
	SQLiteDB = db
	checkSQLiteErr(err)
}

// CreateTable 创建数据库表
func CreateTable(sql string) {
	_, error := SQLiteDB.Exec(sql)

	checkSQLiteErr(error)
}

// Update 更新
func Update(autokid int) {
	trashSQL, err := SQLiteDB.Prepare("update blog set is_deleted='Y',last_modified_at=datetime() where id=?")
	if err != nil {
		checkSQLiteErr(err)
	}
	tx, err := SQLiteDB.Begin()
	if err != nil {
		checkSQLiteErr(err)
	}
	_, err = tx.Stmt(trashSQL).Exec(autokid)
	if err != nil {
		fmt.Println("doing rollback")
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func checkSQLiteErr(err error) {
	if err != nil {
		panic(err)
	}
}
