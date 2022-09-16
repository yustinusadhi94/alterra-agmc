package controllers

import (
	"day-4/lib/database"
	"day-4/middlewares"
	"day-4/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func GetUserById(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := database.GetUserById(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

func CreateUser(c echo.Context) error {
	var data models.User
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := database.CreateUser(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "created",
		"data":   user,
	})
}

func UpdateUser(c echo.Context) error {
	userIdFromToken := middlewares.ExtractTokenUserId(c)

	var data models.User
	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if id != userIdFromToken {
		return c.String(http.StatusUnauthorized, "not allowed to change other user id")
	}

	if err := c.Validate(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := database.UpdateUser(uint(id), data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
		"data":   user,
	})
}

func DeleteUser(c echo.Context) error {
	userIdFromToken := middlewares.ExtractTokenUserId(c)

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if id != userIdFromToken {
		return c.String(http.StatusUnauthorized, "not allowed to change other user id")
	}
	err = database.DeleteUser(uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
