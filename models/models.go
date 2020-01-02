package models

import (
	"database/sql"
	"fmt"
	"log"
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
	TableName  string
	QueryFiled SelectValues
}

// WhereCondition where条件
type WhereCondition struct {
	Operator string
	Value    string
}

// OrderCondition order 条件
type OrderCondition struct {
	OrderBy string
}

// WhereValues where条件值
type WhereValues map[string]WhereCondition

// OrderValues order by 条件值
type OrderValues map[string]OrderCondition

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
	updateString, args := update.MergeUpdate()
	whereString, whereValue := where.mergeWhere()

	for _, v := range whereValue {
		args = append(args, v)
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", p.TableName, updateString, whereString)

	fmt.Println(sql)
	fmt.Println(args)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

// Delete 指定条件的数据
func (p *ModelProperty) Delete(where WhereValues) (int64, error) {
	whereString, whereValue := where.mergeWhere()

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s", p.TableName, whereString)
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(whereValue...)
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
func (p *ModelProperty) Query(
	s SelectValues,
	where WhereValues,
	order OrderValues,
	offset int64,
	limit int64) ([]SelectResult, error) {
	var selectString = s.mergeSelect()

	whereString, whereValue := where.mergeWhere()

	orderBy := order.MergeOrder()

	sql := fmt.Sprintf("SELECT %s FROM %s "+
		"WHERE %s %s LIMIT %d, %d",
		selectString,
		p.TableName,
		whereString,
		orderBy,
		offset,
		limit)

	log.Println(sql)

	rows, err := Conn.Query(sql, whereValue...)

	result := []SelectResult{}

	if err != nil {
		return result, err
	}

	selectField := s.mergeSelectValue()

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return result, err
		}

		tmpResult := s.mergeResultValues()

		result = append(result, tmpResult)
	}

	return result, nil
}

// QueryOne 获取一条数据
func (p *ModelProperty) QueryOne(ele []string, where WhereValues) (SelectResult, error) {
	s := p.QueryFiled
	s = s.filterSelect(ele)

	var selectString = s.mergeSelect()

	whereString, whereValue := where.mergeWhere()

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 0, 1", selectString, p.TableName, whereString)

	rows, err := Conn.Query(sql, whereValue...)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	selectField := s.mergeSelectValue()

	result := []SelectResult{}

	for rows.Next() {
		err = rows.Scan(selectField...)

		if err != nil {
			return nil, err
		}

		tmpResult := s.mergeResultValues()

		result = append(result, tmpResult)
	}

	return result[0], nil
}

// SortedKeys OrderValues
func (o OrderValues) SortedKeys() []string {
	sortedKeys := make([]string, 0)
	for k := range o {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	return sortedKeys
}

// MergeOrder 合并order by条件值
func (o OrderValues) MergeOrder() string {
	order := []string{}

	if len(o) == 0 {
		return ""
	}

	sortedKeys := o.SortedKeys()

	var j = 0
	for _, k := range sortedKeys {
		v := o[k]

		s := fmt.Sprintf("%s %s", k, v.OrderBy)
		order = append(order, s)

		j++
	}

	return " ORDER BY " + strings.Join(order, " , ")
}

// SortedKeys WhereValues
func (w WhereValues) SortedKeys() []string {
	sortedKeys := make([]string, 0)
	for k := range w {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	return sortedKeys
}

func (w WhereValues) mergeWhere() (string, []interface{}) {
	where := []string{}
	value := make([]interface{}, len(w))

	if len(w) == 0 {
		return "1=1", nil
	}

	sortedKeys := w.SortedKeys()

	var j = 0
	for _, k := range sortedKeys {

		v := w[k]
		if v.Operator == "" {
			s := fmt.Sprintf("%s = ?", k)
			where = append(where, s)
		} else {
			s := fmt.Sprintf("%s %s ?", k, v.Operator)
			where = append(where, s)
		}

		value[j] = v.Value
		j++
	}

	return strings.Join(where, " AND "), value
}

// SortedKeys UpdateValues
func (u UpdateValues) SortedKeys() []string {
	sortedKeys := make([]string, 0)
	for k := range u {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	return sortedKeys
}

// MergeUpdate 合并update条件
func (u UpdateValues) MergeUpdate() (string, []interface{}) {
	sortedKeys := u.SortedKeys()

	update := []string{}
	value := make([]interface{}, len(u))

	var j = 0
	for _, k := range sortedKeys {
		s := fmt.Sprintf("%s = ?", k)
		update = append(update, s)
		value[j] = u[k]
		j++
	}

	return strings.Join(update, ", "), value
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

func contains(arr []string, val string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

// FilterSelect 过滤需要的
func (s SelectValues) filterSelect(ele []string) SelectValues {
	if ele != nil {
		for k := range s {
			if !contains(ele, k) {
				delete(s, k)
			}
		}
	}

	return s
}

func (s SelectValues) mergeSelect() string {
	sortedKeys := s.SortSelect()

	value := []string{}
	for _, k := range sortedKeys {
		v := fmt.Sprintf("`%s`", k)
		value = append(value, v)
	}

	return strings.Join(value, ", ")
}

func (s SelectValues) mergeSelectValue() []interface{} {
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

func (s SelectValues) mergeResultValues() SelectResult {
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

// ToString 将指定键值的值转为字符串
func (s SelectResult) ToString(key string) string {

	for _, v := range s {
		var ref = reflect.ValueOf(v)
		var refv = ref.Elem()

		if refv.Kind() == reflect.String {
			fmt.Println(refv.String())
		}
	}

	return ""
	// return strings.Join(arr, splitStr)
}
