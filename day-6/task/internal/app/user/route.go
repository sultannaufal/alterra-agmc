package user

import (
	"example.com/architecture/internal/middleware"
	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group) {
	g.GET("", Get, middleware.AuthMiddleware)
	g.GET("/:id", GetByID, middleware.AuthMiddleware)
	g.POST("", Create)
	g.PUT("/:id", Update, middleware.AuthMiddleware)
	g.DELETE("/:id", Delete, middleware.AuthMiddleware)
}
