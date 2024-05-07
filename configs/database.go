package configs

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(username string, password string, dbName string) *sql.DB {
	poolDB, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		return nil
	}

	// Check if the connection to database is alive
	// If user inserted wrong password, username, or database name, err != nil
	err = poolDB.Ping()
	if err != nil {
		fmt.Println("Failed to connect")
		return nil
	}

	poolDB.SetMaxIdleConns(10)
	poolDB.SetMaxOpenConns(100)
	poolDB.SetConnMaxIdleTime(3 * time.Minute)
	poolDB.SetConnMaxLifetime(60 * time.Minute)
	fmt.Println("Success make connection")
	return poolDB
}
