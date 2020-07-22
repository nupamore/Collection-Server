package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	member struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	members = map[int]*member{
		123: &member{
			ID: 123,
		},
	}
)

func (router) getMember(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, members[id])
}
