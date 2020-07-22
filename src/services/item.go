package services

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ItemService : Items services
type ItemService struct{}

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

// GetItem : get a item
func (ItemService) GetItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, items[id])
}
