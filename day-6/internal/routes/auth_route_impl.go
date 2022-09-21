package routes

import (
	"day-6/internal/entities"
	"day-6/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthRouteImpl struct {
	authService services.AuthService
}

func NewAuthRoute(service services.AuthService) AuthRoute {
	return &AuthRouteImpl{
		authService: service,
	}
}

func (a AuthRouteImpl) AddAuthRoute(e *echo.Echo) *echo.Echo {
	e.POST("/v1/login", a.Login)

	return e
}

func (a AuthRouteImpl) Login(c echo.Context) error {
	var user entities.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := a.authService.Login(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"token":  token,
	})
}
