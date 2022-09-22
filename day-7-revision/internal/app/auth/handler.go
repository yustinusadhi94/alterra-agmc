package auth

import (
	"day-7-revision/internal/factory"
	"day-7-revision/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

type Handler interface {
	Login(c echo.Context) error
}

func (h *handler) Login(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Login(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"token":  token,
	})
}
