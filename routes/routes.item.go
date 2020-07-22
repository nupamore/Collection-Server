package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	items = map[int]*item{
		123: &item{
			ID:   123,
			Name: "test",
		},
	}
)

func (router) getItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, items[id])
}
