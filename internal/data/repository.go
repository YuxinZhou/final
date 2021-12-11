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

type repositoryImpl struct {
	database *sql.DB
}

func GetRepository() biz.Repository {
	once.Do(func() {
		NewRepositoryImpl(db.GetSQLDB())
	})
	return repository
}

func NewRepositoryImpl(dqlDB *sql.DB) biz.Repository {
	if err := db.GetDb().Table("user").AutoMigrate(&biz.User{}); err != nil {
		log.Fatal(err)
	}



	repository = &repositoryImpl{database: dqlDB}
	return repository
}
