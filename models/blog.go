package models

import (
	"database/sql"

	"github.com/durban89/wiki/db"
	// Register MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Conn 连接
var Conn *sql.DB

func init() {
	// Conn = db.DB
	Conn = db.SQLiteDB
}

// Query 获取一条数据
func Query() ([]string, error) {
	rows, err := Conn.Query("SELECT * FROM blog")

	if err != nil {
		return nil, err
	}

	var res = []string{}

	for rows.Next() {
		var autokid int
		var title string
		err = rows.Scan(&autokid, &title)

		if err != nil {
			return nil, err
		}
		res = append(res, title)
	}

	return res, nil
}
