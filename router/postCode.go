package router

import (
	"fmt"
	"net/http"

	"github.com/hackathon-22-spring-16/backend/model"
	"github.com/labstack/echo/v4"
)

type postCodeRequest struct {
	PlainCode string `json:"plainCode"`
	Stdin     string `json:"stdin"`
	Title     string `json:"title"`
	Options   string `json:"options"`
}

type postCodeResponse struct {
	UserName string `json:"userName"`
	Hash     string `json:"hash"`
}

func PostCodeHandler(c echo.Context) error {
	data := &postCodeRequest{}
	err := c.Bind(data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%+v", data))
	}

	userNames, ok := c.Request().Header["X-Showcase-User"]
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get header:X-Showcase-User")
	}
	userName := userNames[0]

	code := model.Code{
		UserName:  userName,
		PlainCode: data.PlainCode,
		Stdin:     data.Stdin,
		Title:     data.Title,
		Options:   data.Options,
	}

	hash, err := model.AddCode(c.Request().Context(), &code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("%w", err))
	}

	res := postCodeResponse{
		UserName: userName,
		Hash:     hash,
	}

	return c.JSON(http.StatusOK, res)
}
