package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbClient *sql.DB
)

func before() error {
	driver := "sqlite3"
	dataSource := ":memory:"
	var err error
	dbClient, err = sql.Open(driver, dataSource)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		os.Exit(-1)
	}

	_, err = dbClient.Exec("CREATE TABLE `test` ( id INT, vv TEXT )")
	if err != nil {
		fmt.Printf("error creating test schema: %v\n", err)
		os.Exit(-2)
	}

	_, err = dbClient.Exec("insert into `test` ( id, vv ) values (1, \"222222\")")
	if err != nil {
		fmt.Printf("error creating test schema: %v\n", err)
		os.Exit(-2)
	}
	return nil
}

func main() {
	err := before()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	var test testStruct

	where := map[string]interface{}{
		"id": 1,
	}

	err = GetOne(dbClient, &test, "test", where)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("%v\n", test.ID)
	fmt.Printf("%v\n", test.VV)
}
