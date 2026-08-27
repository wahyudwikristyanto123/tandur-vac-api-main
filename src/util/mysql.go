package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *sql.DB

func OpenMySQL() {
	if mysql != nil {
		return
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1:3306"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "dev_tandur"
	}

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true&charset=utf8mb4&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open("mysql", uri)
	if err != nil {
		log.Fatalf("[MySQL] Failed to open database: %v", err)
	}

	// Connection Pool Settings for high-performance concurrent throughput
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Printf("[MySQL Warning] Could not ping database at %s: %v", dbHost, err)
	} else {
		log.Println("[MySQL] Database connection pool initialized successfully")
	}

	mysql = db
}

func GetMySQL() *sql.DB {
	return mysql
}

func CloseMySQL() {
	if mysql != nil {
		_ = mysql.Close()
		mysql = nil
	}
}
