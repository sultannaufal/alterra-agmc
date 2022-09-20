package book

import (
	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group) {
	g.GET("", Get)
	// g.GET("/:id", GetByID)
	// g.POST("", Create)
	// g.PUT("/:id", Update, middleware.AuthMiddleware)
	// g.DELETE("/:id", Delete, middleware.AuthMiddleware)
}
