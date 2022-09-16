package routes

import (
	c "example.com/test/controllers"
	"example.com/test/lib/helpers"
	"example.com/test/middlewares"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	books := v1.Group("/books")

	e.Use(middlewares.LogMiddleware)
	e.HTTPErrorHandler = helpers.ErrorHandler
	e.Validator = helpers.Validator

	users.GET("", c.GetUsersController, middlewares.AuthMiddleware)
	users.GET("/:id", c.GetUserByIDController, middlewares.AuthMiddleware)
	users.POST("", c.AddUserController)
	users.POST("/token", c.GenerateTokenController)
	users.PUT("/:id", c.UpdateUserController, middlewares.AuthMiddleware)
	users.DELETE("/:id", c.DeleteUserController, middlewares.AuthMiddleware)

	books.GET("", c.GetBooksController)
	books.GET("/:id", c.GetBookByIDController)
	books.POST("", c.AddBookController, middlewares.AuthMiddleware)
	books.PUT("/:id", c.UpdateBookController, middlewares.AuthMiddleware)
	books.DELETE("/:id", c.DeleteBookController, middlewares.AuthMiddleware)

	return e
}
