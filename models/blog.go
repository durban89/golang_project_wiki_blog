package models

import (
	"database/sql"
	"fmt"

	"github.com/durban89/wiki/db"
	"github.com/durban89/wiki/helpers"
)

var tableName = "blog"

// Conn 连接
var Conn *sql.DB

func init() {
	// MySQL
	Conn = db.MySQLDB

	// SQLite
	// Conn = db.SQLiteDB

	// PostgreSQL
	// Conn = db.PostgreSQLDB
}

// Create 添加数据
func Create(p *helpers.Page) (int64, error) {
	sql := fmt.Sprintf("INSERT %s SET title=?", tableName)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Title)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update 更新数据
func Update(p *helpers.Page, id int64) (int64, error) {
	sql := fmt.Sprintf("UPDATE %s SET title=? WHERE autokid=?", tableName)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Title, id)
	if err != nil {
		return 0, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

// Delete 删除数据
func Delete(id int64) (int64, error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE autokid=?", tableName)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

// Query 获取数据
func Query() ([]helpers.Page, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := Conn.Query(sql)

	if err != nil {
		return nil, err
	}

	var res = []helpers.Page{}

	for rows.Next() {
		var autokid int
		var title string
		err = rows.Scan(&autokid, &title)

		if err != nil {
			return nil, err
		}

		p := helpers.Page{
			ID:    autokid,
			Title: title,
		}

		res = append(res, p)
	}

	return res, nil
}

// QueryOne 获取一条数据
func QueryOne() ([]helpers.Page, error) {
	sql := fmt.Sprintf("SELECT * FROM %s LIMIT 0, 1", tableName)

	rows, err := Conn.Query(sql)

	if err != nil {
		return nil, err
	}

	var res = []helpers.Page{}

	for rows.Next() {
		var autokid int
		var title string
		err = rows.Scan(&autokid, &title)

		if err != nil {
			return nil, err
		}

		p := helpers.Page{
			ID:    autokid,
			Title: title,
		}

		res = append(res, p)
	}

	return res, nil
}
