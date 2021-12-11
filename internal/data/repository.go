package data

import (
	"database/sql"
	"log"
	"sync"

	"final/internal/biz"
	"final/internal/db"
)

// repository singleton
var repository biz.Repository
var once sync.Once

const userTableName = "users"

type repositoryImpl struct {
	db *sql.DB
}

func GetRepository() biz.Repository {
	once.Do(func() {
		NewRepositoryImpl(db.GetSQLDB())
	})
	return repository
}

func NewRepositoryImpl(dqlDB *sql.DB) biz.Repository {
	// initialize DB schema
	if err := db.GetDb().Table(userTableName).AutoMigrate(&biz.User{}); err != nil {
		log.Fatal(err)
	}

	repository = &repositoryImpl{db: dqlDB}
	return repository
}
