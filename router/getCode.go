package router

import (
	"net/http"

	"github.com/hackathon-22-spring-16/backend/model"
	"github.com/labstack/echo/v4"
)

// GET /get-code/:username/:hash
func GetCodeHandler(c echo.Context) error {
	username := c.Param("username")
	hash := c.Param("hash")

	code, err := model.GetCode(c.Request().Context(), username, hash)

	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	if code == nil {
		return c.String(http.StatusInternalServerError, "GetCode が実装されてないよ")
	}
	return c.JSON(200, &code)
}