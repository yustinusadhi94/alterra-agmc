package routes

import (
	"day-3/controllers"
	"day-3/middlewares"
	"day-3/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	books := e.Group("/v1/books")
	books.GET("", controllers.GetBooks)
	books.GET("/:id", controllers.GetBookById)
	books.POST("", controllers.CreateBook, middleware.JWT([]byte(middlewares.SecretKey)))
	books.PUT("/:id", controllers.UpdateBook, middleware.JWT([]byte(middlewares.SecretKey)))
	books.DELETE("/:id", controllers.DeleteBook, middleware.JWT([]byte(middlewares.SecretKey)))

	users := e.Group("/v1/users")
	users.GET("", controllers.GetUsers, middleware.JWT([]byte(middlewares.SecretKey)))
	users.GET("/:id", controllers.GetUserById, middleware.JWT([]byte(middlewares.SecretKey)))
	users.POST("", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUser, middleware.JWT([]byte(middlewares.SecretKey)))
	users.DELETE("/:id", controllers.DeleteUser, middleware.JWT([]byte(middlewares.SecretKey)))

	e.POST("/v1/login", controllers.LoginUser)

	return e
}
