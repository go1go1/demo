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

	testInject()
}

// 自行拼接sql会导致注入
func queryMulti(name string) {
	//
	sqlStr := fmt.Sprintf("select id, name, age from users where name='%s'", name)
	var users []User

	err := DB.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, v := range users {
		fmt.Printf("%v\n", v)
	}
}

// 使用bind模式
func queryMultiDef(name string) {
	//自行拼接sql
	sqlStr := "select id, name, age from users where name=?"
	var users []User

	err := DB.Select(&users, sqlStr, name)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, v := range users {
		fmt.Printf("%v\n", v)
	}
}

func testInject() {
	//构造一个注入,返回表所有记录
	//queryMulti("tom' or 1=1 #")
	//猜测表的记录条数(有返回值表示表记录小于10)
	//queryMulti("tom' and (select count(*) from users) < 10 #")
	//使用union返回表所有记录
	//queryMulti("tom' union select * from users #")

	//bind模式能够避免被注入
	queryMultiDef("tom' or 1=1 #")
	queryMultiDef("tom' union select * from users #")
}
