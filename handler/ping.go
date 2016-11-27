package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		json := map[string]string{
			"Message": "pong",
		}
		return c.JSON(http.StatusOK, json)
	}
}
