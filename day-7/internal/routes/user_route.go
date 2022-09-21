package routes

import "github.com/labstack/echo/v4"

type UserRoute interface {
	AddUserRoutes(e *echo.Echo, TokenValidator echo.MiddlewareFunc) *echo.Echo
	GetUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
