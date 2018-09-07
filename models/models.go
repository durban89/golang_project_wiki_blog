package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/durban89/wiki/db"
	"github.com/durban89/wiki/helpers"
)

// ModelMethod 接口
type ModelMethod interface {
	Query(offset int64, limit int64) (map[string]string, error)
	Create() (int64, error)
	Delete() (int64, error)
	Update() (int64, error)
	QueryOne(int64) (map[string]string, error)
}

// ModelProperty 属性
type ModelProperty struct {
	TableName string
}

// WhereCondition where条件
type WhereCondition struct {
	Operator string
	Value    string
}

// WhereValues where条件值
type WhereValues map[string]WhereCondition

// UpdateValues update条件值
type UpdateValues map[string]string

// InsertValues Insert值
type InsertValues map[string]string

// SelectValues select条件值
type SelectValues []string

// Conn 连接
var Conn *sql.DB

func init() {
	// MySQL
	Conn = db.MySQLDB

	// SQLite
	// Conn = db.SQLiteDB

	// PostgreSQL
	// Conn = db.PostgreSQLDB

	// var blogInstance BlogModel
	// blogInstance.Where()
	// blogInstance.Update()
}

// Create 添加数据
func (p *ModelProperty) Create(data InsertValues) (int64, error) {
	name, value := data.MergeInsert()
	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s", tableName, name, value)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update 更新
func (p *ModelProperty) Update(update UpdateValues, where WhereValues) (int64, error) {
	var updateString = update.MergeUpdate()
	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, updateString, whereString)

	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec()
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
func (p *ModelProperty) Delete(id int64) (int64, error) {
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
func (p *ModelProperty) Query() ([]helpers.Page, error) {
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
func (p *ModelProperty) QueryOne(s SelectValues, where WhereValues) (helpers.Page, error) {
	var selectString = s.MergeSelect()

	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 0, 1", selectString, tableName, whereString)

	rows, err := Conn.Query(sql)

	var res = helpers.Page{}

	if err != nil {
		return res, err
	}

	var autokid int
	var title string

	for rows.Next() {
		err = rows.Scan(&autokid, &title)
		// err = rows.Scan(interfaceSlice...)

		if err != nil {
			return res, err
		}

		p := helpers.Page{
			ID:    autokid,
			Title: title,
		}

		res = p
	}

	if res.ID == 0 {
		return res, errors.New("资源不存在")
	}

	return res, nil
}

// MergeWhere 合并where条件
func (w WhereValues) MergeWhere() string {
	where := []string{}
	for k, v := range w {
		if v.Operator == "" {
			s := fmt.Sprintf("%s = %s", k, v.Value)
			where = append(where, s)
		} else {
			s := fmt.Sprintf("%s %s %s", k, v.Operator, v.Value)
			where = append(where, s)
		}
	}

	return strings.Join(where, " AND ")
}

// MergeUpdate 合并update条件
func (u UpdateValues) MergeUpdate() string {
	update := []string{}
	for k, v := range u {
		s := fmt.Sprintf("%s = '%s'", k, v)
		update = append(update, s)
	}

	return strings.Join(update, ", ")
}

// MergeInsert 合并Insert值
func (i InsertValues) MergeInsert() (string, string) {
	name := []string{}
	value := []string{}
	for k, v := range i {
		n := fmt.Sprintf("%s", k)
		v := fmt.Sprintf("'%s'", v)
		name = append(name, n)
		value = append(value, v)
	}

	return strings.Join(name, ", "), fmt.Sprintf("(%s)", strings.Join(value, ", "))
}

// MergeSelect  合并select条件
func (s SelectValues) MergeSelect() string {
	return strings.Join(s, ", ")
}
