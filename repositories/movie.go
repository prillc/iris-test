package repositories

import (
	"iris-test/models"
	"sync"
)

type MovieRepository interface {
	Select() models.Movie
}

func NewMovieRepository() *movieRepository {
	return &movieRepository{}
}

type movieRepository struct {
	mu sync.Mutex
	// db gorm.DB
}

func (m *movieRepository) Select() models.Movie {
	return models.Movie{ID: 1, Name: "golang web", Year: 2021, Genre: "perfect", Poster: "SanZ"}
}
