package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gDb *gorm.DB
var sqlDB *sql.DB

const maxDBConnections = 3

// InitDb initializes the global handle.
func InitDb() {
	dbName := "user_system"
	connStr := fmt.Sprintf("root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName) // connect to local DB

	var err error
	gDb, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err = gDb.DB()
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDb closes the database handle.
func CloseDb() {
	err := sqlDB.Close()
	if err != nil {
		log.Printf("Failed to close the database handle: %v", err)
	}
}

// GetDb returns the global handle to database.
func GetDb() *gorm.DB {
	return gDb
}

// SetSQLDB is primarily useful for setting up a mock db implementation during testing.
func SetSQLDB(db *sql.DB) {
	sqlDB = db
}

// GetSQLDB returns global sql.db handle.
func GetSQLDB() *sql.DB {
	return sqlDB
}