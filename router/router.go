package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://front-end.hackathon-22-spring-16.trap.show"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
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