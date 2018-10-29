package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

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

// ModelError 模型错误信息结构
type ModelError struct {
	When time.Time
	What string
}

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

func (e ModelError) Error() string {
	return fmt.Sprintf("时间 %v, 错误信息%s", e.When, e.What)
}

// Create 添加数据
func (p *ModelProperty) Create(data InsertValues) (int64, error) {
	name, preValue, value := data.MergeInsert()
	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s", p.TableName, name, preValue)

	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(value...)
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

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", p.TableName, updateString, whereString)

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
	sql := fmt.Sprintf("DELETE FROM %s WHERE autokid=?", p.TableName)
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
		selectString, p.TableName, whereString, offset, limit)

	rows, err := Conn.Query(sql)

	result := []SelectResult{}

	if err != nil {
		return result, err
	}

	selectField := s.MergeSelectValue()

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return result, err
		}

		tmpResult := s.MergeResultValues()

		result = append(result, tmpResult)
	}

	return result, nil
}

// QueryOne 获取一条数据
func (p *ModelProperty) QueryOne(s SelectValues, where WhereValues) error {
	var selectString = s.MergeSelect()

	var whereString = where.MergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 0, 1", selectString, p.TableName, whereString)

	rows, err := Conn.Query(sql)

	if err != nil {
		return err
	}

	// columns, err := rows.Columns()
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(columns)

	// columnTypes, err := rows.ColumnTypes()
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(columnTypes)

	selectField := s.MergeSelectValue()

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return err
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
func (i InsertValues) MergeInsert() (string, string, []interface{}) {
	sortedKeys := i.SortSelect()

	name := []string{}
	preValue := []string{}
	value := make([]interface{}, len(i))
	var j = 0
	for _, k := range sortedKeys {
		n := fmt.Sprintf("%s", k)
		v := fmt.Sprintf("?")
		name = append(name, n)
		preValue = append(preValue, v)
		fmt.Println("k = ", k)
		fmt.Println("v = ", i[k])
		value[j] = i[k]
		j++
	}

	return fmt.Sprintf("(%s)", strings.Join(name, ", ")), fmt.Sprintf("(%s)", strings.Join(preValue, ", ")), value
}

// MergeSelect  合并select条件
func (s SelectValues) MergeSelect() string {
	sortedKeys := s.SortSelect()

	value := []string{}
	for _, k := range sortedKeys {
		v := fmt.Sprintf("`%s`", k)
		value = append(value, v)
	}

	return strings.Join(value, ", ")
}

// MergeSelectValue  取出select条件的值
func (s SelectValues) MergeSelectValue() []interface{} {
	sortedKeys := s.SortSelect()

	selectField := make([]interface{}, len(s))
	var i = 0
	for _, k := range sortedKeys {
		selectField[i] = s[k]
		i++
	}

	return selectField
}

// SortSelect 排序select keys
func (s SelectValues) SortSelect() []string {
	sortedKeys := make([]string, 0)
	for k := range s {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	return sortedKeys
}

// SortSelect 排序insert keys
func (i InsertValues) SortSelect() []string {
	sortedKeys := make([]string, 0)
	for k := range i {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	return sortedKeys
}

// MergeResultValues 查询结果的值合并
func (s SelectValues) MergeResultValues() SelectResult {
	var tmpResult = SelectResult{}

	var i = 0

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

	return tmpResult
}
