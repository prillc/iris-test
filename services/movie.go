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

type MovieHistoryService interface {
	List(query interface{}, limit int) []*models.MovieHistory
	Count(query interface{}) int64
}

type movieHistoryService struct {
	repo repositories.MovieHistoryRepository
}

func NewMovieHistoryService(repo repositories.MovieHistoryRepository) MovieHistoryService {
	return &movieHistoryService{repo: repo}
}

func (s *movieHistoryService) List(query interface{}, limit int) []*models.MovieHistory {
	return s.repo.Select(query, limit)
}

func (s *movieHistoryService) Count(query interface{}) int64 {
	return s.repo.Count(query)
}
