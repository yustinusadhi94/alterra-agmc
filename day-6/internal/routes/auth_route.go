package routes

import "github.com/labstack/echo/v4"

type AuthRoute interface {
	AddAuthRoute(e *echo.Echo) *echo.Echo
	Login(c echo.Context) error
}
