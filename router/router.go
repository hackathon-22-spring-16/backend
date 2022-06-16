package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SetRouting() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	e.Start(port)
}