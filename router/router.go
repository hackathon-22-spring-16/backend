package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func SetRouting() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/echo", func(c echo.Context) error {
		header := fmt.Sprintf("%#v", c.Request().Header)
		return c.String(http.StatusOK, header)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	e.Start(port)
}