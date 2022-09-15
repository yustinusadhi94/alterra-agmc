package controllers

import (
	"day-3/lib/database"
	"day-3/models"
	"github.com/labstack/echo"
	"net/http"
)

func LoginUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := database.LoginUser(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"token":  token,
	})
}
