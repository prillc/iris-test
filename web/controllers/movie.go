package controllers

import (
	"github.com/kataras/iris/v12"
	"iris-test/common"
	"iris-test/models"
	"iris-test/services"
)

type MovieController struct {
	Service services.MovieService
}

// localhost:8080/movie/ping
func (c *MovieController) GetPing() string {
	return "pong"
}

// localhost:8080/movie
func (c *MovieController) Get() *common.Response {
	//movie, _ := c.Service.GetByID(1)
	return &common.Response{
		Status: 0,
		Msg:    "",
		Data: iris.Map{
			"movies": c.Service.Select(10),
		},
	}
}

// localhost:8080/movie/1
func (c *MovieController) GetBy(id int64) *common.Response {
	return &common.Response{
		Status: 0,
		Msg:    "",
		Data:   c.Service.FindById(id),
	}
}

// // localhost:8080/movie
func (c *MovieController) Post(ctx iris.Context) *common.Response {
	// 处理参数让后返回
	return common.SuccessResponse(models.Movie{
		ID:     2,
		Name:   "",
		Year:   2021,
		Genre:  "",
		Poster: "",
	}, "ok")
}

type MovieHistoryController struct {
	Service services.MovieHistoryService
}

func (m *MovieHistoryController) Get() *common.Response {
	histories := m.Service.List(map[string]interface{}{"movie_id": 1}, 10)
	return &common.Response{
		Status: 0,
		Msg:    "",
		Data: iris.Map{
			"histories": histories,
		},
	}
}

func (m *MovieHistoryController) GetCount() common.Response {
	count := m.Service.Count(map[string]interface{}{"movie_id": 1})
	return common.Response{
		Status: 0,
		Msg:    "",
		Data: iris.Map{
			"count": count,
		},
	}
}
