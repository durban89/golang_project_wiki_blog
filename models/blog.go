package models

import (
	"github.com/durban89/wiki/db"
	// Register MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Query 获取一条数据
func Query() ([]string, error) {
	rows, err := db.DB.Query("SELECT * FROM blog")

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
