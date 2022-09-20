package middleware

import (
	"example.com/architecture/pkg/util/log"
	"github.com/labstack/echo/v4"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}
