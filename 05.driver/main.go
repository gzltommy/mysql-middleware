package main

import (
	"database/sql"

	_ "github.com/go-mysql-org/go-mysql/driver"
)

func main() {
	// dsn format: "user:password@addr?dbname"
	dsn := "root@127.0.0.1:3306?test"
	db, _ := sql.Open("mysql", dsn)
	db.Close()
}
