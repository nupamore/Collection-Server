package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/labstack/echo/v4"
)

// Connect : Database connect
func Connect() *gorm.DB {
	dbTargetStr := fmt.Sprintf("%s:%s@(%s)/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		"collection",
		"charset=utf8&parseTime=True&loc=Local",
	)
	db, err := gorm.Open("mysql", dbTargetStr)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	return db
}

// ContextDB : inject echo context
func ContextDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
