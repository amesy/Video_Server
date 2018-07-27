package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("mysql", "mysqlcli:12345678@tcp(10.68.7.24:3306)/video_server?charset=utf8")

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}
