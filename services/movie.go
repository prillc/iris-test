package services

import (
	"iris-test/models"
	"iris-test/repositories"
)

type MovieService interface {
	GetByID(id int64) (models.Movie, bool)
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

type movieService struct {
	repo repositories.MovieRepository
}

func (s *movieService) GetByID(id int64) (models.Movie, bool) {
	return s.repo.Select(), true
}
