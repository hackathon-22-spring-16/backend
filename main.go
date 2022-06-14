package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/hackathon-22-spring-16/backend/model"
)

func main() {
	db, err := model.InitDB()

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/db-info", func(c echo.Context) error {
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("db error: %s", err))
		}
		return c.String(http.StatusOK, fmt.Sprintf("driver name: %s", db.DriverName()))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	e.Start(port)
}
