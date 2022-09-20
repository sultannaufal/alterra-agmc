package http

import (
	"example.com/architecture/internal/app/auth"
	"example.com/architecture/internal/app/book"
	"example.com/architecture/internal/app/user"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")
	user.Route(v1.Group("/users"))
	auth.Route(v1.Group("/auth"))
	book.Route(v1.Group("/books"))
}
