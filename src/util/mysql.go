package util

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysql *sql.DB

// LoadEnv reads key=value pairs from a .env file and sets them if not already present
func LoadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			if os.Getenv(key) == "" {
				_ = os.Setenv(key, val)
			}
		}
	}
}

func OpenMySQL() {
	if mysql != nil {
		return
	}

	LoadEnv(".env")

	dbUser := os.Getenv("DB_USERNAME")
	if dbUser == "" {
		dbUser = os.Getenv("DB_USER")
		if dbUser == "" {
			dbUser = "admin"
		}
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = os.Getenv("DB_PASS")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "149.28.136.226"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" && !strings.Contains(dbHost, ":") {
		dbHost = fmt.Sprintf("%s:%s", dbHost, dbPort)
	} else if !strings.Contains(dbHost, ":") {
		dbHost = fmt.Sprintf("%s:3306", dbHost)
	}

	dbName := os.Getenv("DB_DATABASE")
	if dbName == "" {
		dbName = os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "tandur_vac"
		}
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
		log.Printf("[MySQL] Database '%s' on %s connected successfully", dbName, dbHost)
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
