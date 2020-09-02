package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type User struct {
	Id int `bson:"_id" json:"id"`
	//Name string `bson:"name" json:"name"`
	// 如果mysql字段值允许null，需要使用sql.NullString
	Name sql.NullString `bson:"name" json:"name"`
	Age  int            `bson:"age" json:"age"`
}

func initDB() error {
	var err error
	dsn := "root:123456@tcp(9.135.12.51)/test"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)

	return nil
}

func main() {
	initDB()

	queryRow()

	queryMulti()

	execInsert()
}

// 查询单条
func queryRow() {
	sqlStr := "select id, name, age from users where id=?"
	var user User
	err := DB.Get(&user, sqlStr, 2)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("%#v\n", &user)
}

// 查询多条
func queryMulti() {
	sqlStr := "select id, name, age from users where id>?"
	var users []User

	err := DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%#v\n", &users)
}

// 插入
func execInsert() {
	sqlStr := "insert into users (name, age) values(?,?)"
	res, err := DB.Exec(sqlStr, "tom", 17)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("in is: %d\n", id)
}
