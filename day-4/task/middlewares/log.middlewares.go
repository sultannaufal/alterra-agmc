package middlewares

import (
	"example.com/test/lib/helpers"
	"github.com/labstack/echo/v4"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		helpers.MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}
