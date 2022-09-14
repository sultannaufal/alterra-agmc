package routes

import (
	"example.com/crud/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	api := e.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	books := v1.Group("/books")

	users.GET("", controllers.GetUsersController)
	users.GET("/:id", controllers.GetUserByIDController)
	users.POST("", controllers.AddUserController)
	users.PUT("/:id", controllers.UpdateUserController)
	users.DELETE("/:id", controllers.DeleteUserController)

	books.GET("", controllers.GetBooksController)
	books.GET("/:id", controllers.GetBookByIDController)
	books.POST("", controllers.AddBookController)
	books.PUT("/:id", controllers.UpdateBookController)
	books.DELETE("/:id", controllers.DeleteBookController)

	return e
}
