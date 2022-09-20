package middleware

import (
	"net/http"

	"example.com/architecture/pkg/util/validator"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),
	)
	e.Use(LogMiddleware)
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = validator.Validator

}
