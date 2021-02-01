package repositories

import (
	"gorm.io/gorm"
	"iris-test/models"
	"sync"
)

type MovieRepository interface {
	FindById(id int64) *models.Movie
	Select(limit int) []*models.Movie
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

type movieRepository struct {
	mu sync.Mutex
	db *gorm.DB
}

func (m *movieRepository) FindById(id int64) *models.Movie {
	var movie models.Movie
	m.db.First(&movie, id)
	return &movie
}

func (m *movieRepository) Select(limit int) []*models.Movie {
	var movies []*models.Movie
	m.db.Limit(limit).Order("id desc").Find(&movies)
	return movies
}

type MovieHistoryRepository interface {
	FindById(id int64) *models.MovieHistory
	Select(query interface{}, limit int) []*models.MovieHistory
	Count(query interface{}) int64
}

func NewMovieHistoryRepository(db *gorm.DB) MovieHistoryRepository {
	return &movieHistoryRepository{db: db}
}

type movieHistoryRepository struct {
	mu sync.Mutex
	db *gorm.DB
}

func (m *movieHistoryRepository) FindById(id int64) *models.MovieHistory {
	var movieHistory models.MovieHistory
	m.db.First(&movieHistory, id)
	return &movieHistory
}

func (m *movieHistoryRepository) Select(query interface{}, limit int) []*models.MovieHistory {
	var movieHistories []*models.MovieHistory
	m.db.Limit(limit).Where(query).Find(&movieHistories)
	return movieHistories
}

func (m *movieHistoryRepository) Count(query interface{}) int64 {
	var count int64
	m.db.Model(&models.MovieHistory{}).Find(query).Count(&count)
	return count
}
