package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	fmt.Println("Entering conn.go init function...")
	dbConn, err = sql.Open("mysql",
		"mysqlcli:12345678@tcp(10.68.7.24:3306)/video_server_new?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("dbConn +%v\n", dbConn)
}
