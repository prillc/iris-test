package services

import (
	"iris-test/models"
	"iris-test/repositories"
)

type MovieService interface {
	FindById(id int64) *models.Movie
	Select(limit int) []*models.Movie
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

type movieService struct {
	repo repositories.MovieRepository
}

func (s *movieService) Select(limit int) []*models.Movie {
	return s.repo.Select(limit)
}

func (s *movieService) FindById(id int64) *models.Movie {
	return s.repo.FindById(id)
}
