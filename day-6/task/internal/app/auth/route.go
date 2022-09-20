package auth

import (
	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group) {
	g.POST("/login", Token)
}
