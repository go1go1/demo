package main

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

func main() {
	initDB()
	defer DB.Close()

	//queryRow()
	//queryMulti()
	//execInsert()

	prePareQuery()
	//prePareInsert()
}

func initDB() error {
	dsn := "root:123456@tcp(127.0.0.1)/test"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

// 查询单条
func queryRow() {
	sqlStr := "select id, name, age from users where id=?"
	row := DB.QueryRow(sqlStr, 2)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("%#v\n", &user)
}

// 查询多条
func queryMulti() {
	sqlStr := "select id, name, age from users where id>?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //rows使用后一定要释放
		}
	}()

	var user User
	var users []User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		users = append(users, user)
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

//预处理
func prePareQuery() {
	sqlStr := "select id, name, age from users where id>?"

	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error: %v", err)
		return
	}

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //rows使用后一定要释放
		}
	}()

	var user User
	var users []User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		users = append(users, user)
	}

	fmt.Printf("%#v\n", &users)
}

func prePareInsert() {
	sqlStr := "insert into users (name, age) values(?,?)"
	stmt, err := DB.Prepare(sqlStr)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	res, err := stmt.Exec("tom", 17)
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
