package mysql

import (
	"database/sql"
	_ "embed"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// go:embed ../../.env
var env string

func GetConnection() *sql.DB {
	// if err := godotenv.Load(env); err != nil {
	// 	log.Println("No .env file found")
	// 	panic(err)
	// }

	// var DB_NAME string = os.Getenv("MYSQL_DB_NAME")
	// var DB_PASS string = os.Getenv("DB_PASS")
	// var DB_HOST string = os.Getenv("DB_HOST")
	// var MYSQL_DB_PORT string = os.Getenv("MYSQL_DB_PORT")

	var URI string = "root:dompetkilat@tcp(localhost:3306)/DompetKilatDB"

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
