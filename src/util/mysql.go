package util

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *sql.DB

func OpenMySQL() {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// defer db.Close()
	mysql = db
}

func GetMySQL() *sql.DB {
	return mysql
}

func CloseMySQL() {
	mysql.Close()
}
