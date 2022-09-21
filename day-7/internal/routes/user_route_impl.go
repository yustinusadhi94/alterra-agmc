package routes

import (
	"day-7/internal/entities"
	"day-7/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserRouteImpl struct {
	userService services.UserService
	validator   echo.Validator
}

func NewUserRoute(service services.UserService) UserRoute {
	return &UserRouteImpl{
		userService: service,
	}
}

func (u UserRouteImpl) AddUserRoutes(e *echo.Echo, tokenValidator echo.MiddlewareFunc) *echo.Echo {
	users := e.Group("/v1/users")
	users.GET("", u.GetUsers, tokenValidator)
	users.GET("/:id", u.GetUserById, tokenValidator)
	users.POST("", u.CreateUser)
	users.PUT("/:id", u.UpdateUser, tokenValidator)
	users.DELETE("/:id", u.DeleteUser, tokenValidator)

	return e
}

func (u UserRouteImpl) GetUsers(c echo.Context) error {
	users, err := u.userService.GetUsers()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func (u UserRouteImpl) GetUserById(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := u.userService.GetUserById(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

func (u UserRouteImpl) CreateUser(c echo.Context) error {
	var data entities.UserCreateRequest
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := u.userService.CreateUser(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "created",
		"data":   user,
	})
}

func (u UserRouteImpl) UpdateUser(c echo.Context) error {
	var data entities.UserUpdateRequest
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	//if id != userIdFromToken {
	//	return c.String(http.StatusUnauthorized, "not allowed to change other user id")
	//}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := u.userService.UpdateUser(uint(id), data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
		"data":   user,
	})
}

func (u UserRouteImpl) DeleteUser(c echo.Context) error {
	//userIdFromToken := middlewares.ExtractTokenUserId(c)

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	//if id != userIdFromToken {
	//	return c.String(http.StatusUnauthorized, "not allowed to change other user id")
	//}
	err = u.userService.DeleteUser(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
