package user

import (
	"day-7-revision/internal/factory"
	"day-7-revision/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	GetUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func (h *handler) GetUsers(c echo.Context) error {
	result, err := h.service.GetUsers()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *handler) GetUserById(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.GetUserById(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *handler) CreateUser(c echo.Context) error {
	var data model.UserCreateRequest
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateUser(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "created",
		"data":   result,
	})
}

func (h *handler) UpdateUser(c echo.Context) error {
	var data model.UserUpdateRequest
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.UpdateUser(uint(id), data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
		"data":   user,
	})
}

func (h *handler) DeleteUser(c echo.Context) error {
	//userIdFromToken := middlewares.ExtractTokenUserId(c)

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
