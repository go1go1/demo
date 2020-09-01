package main

import (
	"database/sql"
	"errors"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
)

func init() {
	scanner.SetTagName("json")
}

//GetOne gets one record from table by condition "where"
func GetOne(db *sql.DB, entity interface{}, table string, where map[string]interface{}, selectFields ...[]string) error {
	if "" == table {
		return errors.New("table name couldn't be empty")
	}
	if nil == db {
		return errors.New("sql.DB object couldn't be nil")
	}
	var field []string
	if len(selectFields) > 0 {
		field = selectFields[0]
	} else {
		field = nil
	}
	// limit
	if where == nil {
		where = map[string]interface{}{}
	}
	where["_limit"] = []uint{0, 1}
	cond, val, err := builder.BuildSelect(table, where, field)
	if nil != err {
		return err
	}
	row, err := db.Query(cond, val...)
	if nil != err || nil == row {
		return err
	}
	defer row.Close()
	return scanner.Scan(row, entity)
}
