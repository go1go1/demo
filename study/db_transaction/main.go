package main

/**
mysql 事务
*/
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	Id int `bson:"_id" json:"id"`
	//Name string `bson:"name" json:"name"`
	// 如果mysql字段值允许null，需要使用sql.NullString
	Name sql.NullString `bson:"name" json:"name"`
	Age  int            `bson:"age" json:"age"`
}

func initDB() error {
	dsn := "root:123456@tcp(9.135.12.51)/test"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	initDB()
	defer DB.Close()

	trans()

}

func trans() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("transaction failed, err: %v\n", err)
		return
	}

	sqlStr := "update users set age=2 where id = ?"
	_, err = conn.Exec(sqlStr, 2)
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("exec sql: %s faild, error: %v", sqlStr, err)
		return
	}
	sqlStr2 := "update users set age=3; where id = ?"
	_, err = conn.Exec(sqlStr2, 3)
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Printf("exec sql: %s faild, error: %v", sqlStr2, err)
		return
	}

	err = conn.Commit()
	if err != nil {
		fmt.Printf("commit failed, err: %v\n", err)
		if conn != nil {
			conn.Rollback()
		}
		return
	}
}
