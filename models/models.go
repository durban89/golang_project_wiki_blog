package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/durban89/wiki/db"
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
type SelectValues map[string]interface{}

// SelectResult 结果值
type SelectResult map[string]interface{}

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
func (p *ModelProperty) Query(s SelectValues, where WhereValues, offset int64, limit int64) ([]SelectResult, error) {
	var selectString = s.MergeSelect()

	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT %d, %d",
		selectString, tableName, whereString, offset, limit)

	rows, err := Conn.Query(sql)

	result := []SelectResult{}

	if err != nil {
		return result, err
	}

	selectField := make([]interface{}, len(s))

	var i = 0
	for _, v := range s {
		selectField[i] = v
		i++
	}

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return result, err
		}

		var i = 0
		var tmpResult = SelectResult{}

		for k, v := range s {
			var ref = reflect.ValueOf(v)
			var refv = ref.Elem()

			if refv.Kind() == reflect.Int64 {
				tmpResult[k] = refv.Int()
			} else if refv.Kind() == reflect.String {
				tmpResult[k] = refv.String()
			}

			i++
		}

		result = append(result, tmpResult)
	}

	return result, nil
}

// QueryOne 获取一条数据
func (p *ModelProperty) QueryOne(s SelectValues, where WhereValues) error {
	var selectString = s.MergeSelect()

	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 0, 1", selectString, tableName, whereString)

	rows, err := Conn.Query(sql)

	if err != nil {
		return err
	}

	selectField := make([]interface{}, len(s))

	var i = 0
	for _, v := range s {
		selectField[i] = v
		i++
	}

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return err
		}

		var i = 0
		for _, v := range s {
			ref := reflect.ValueOf(v)
			fmt.Println(ref.Elem())
			i++
		}
	}

	return nil
}

// MergeWhere 合并where条件
func (w WhereValues) MergeWhere() string {
	where := []string{}
	if len(w) == 0 {
		return "1=1"
	}

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
	value := []string{}
	for k := range s {
		v := fmt.Sprintf("`%s`", k)
		value = append(value, v)
	}

	return strings.Join(value, ", ")
}
