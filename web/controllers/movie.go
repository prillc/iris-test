package controllers

import (
	"iris-test/models"
	"iris-test/services"
)

type MovieController struct {
	Service services.MovieService
}

func (c *MovieController) Get() models.Movie {
	movie, _ := c.Service.GetByID(1)
	return movie
}
