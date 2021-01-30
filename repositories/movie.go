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

func NewMovieRepository(db *gorm.DB) *movieRepository {
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
