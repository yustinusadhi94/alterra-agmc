package routes

import (
	"day-2/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	books := e.Group("/v1/books")
	books.GET("", controllers.GetBooks)
	books.GET("/:id", controllers.GetBookById)
	books.POST("", controllers.CreateBook)
	books.PUT("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)

	users := e.Group("/v1/users")
	users.GET("", controllers.GetUsers)
	users.GET("/:id", controllers.GetUserById)
	users.POST("", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)

	return e
}
