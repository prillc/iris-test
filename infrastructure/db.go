package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"iris-test/repositories"
	"log"
)

type Repositories struct {
	DB                     *gorm.DB
	MovieRepository        repositories.MovieRepository
	MovieHistoryRepository repositories.MovieHistoryRepository
}

func NewRepositories() *Repositories {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库打开失败!")
		return nil
	}
	return &Repositories{
		DB:                     db,
		MovieRepository:        repositories.NewMovieRepository(db),
		MovieHistoryRepository: repositories.NewMovieHistoryRepository(db),
	}
}

// 全局的repositories对象
var GlobalRepositories = NewRepositories()
