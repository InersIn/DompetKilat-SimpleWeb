package mysql

import (
	"database/sql"
	_ "embed"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// go:embed ../../.env
var env string

func GetConnection() *sql.DB {

	var DB_NAME string = os.Getenv("MYSQL_DATABASE")
	var DB_PASS string = os.Getenv("MYSQL_PASSWORD")
	var DB_HOST string = os.Getenv("MYSQL_HOST")
	var DB_PORT string = os.Getenv("MYSQL_DB_PORT")

	var URI string = "root:"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME

	db, err := sql.Open("mysql", URI)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
