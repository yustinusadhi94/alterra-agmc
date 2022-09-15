package controllers

import (
	"day-3/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func GetBooks(c echo.Context) error {
	var books []models.Book

	books = append(books, models.Book{
		Id:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Test Title",
		Author:      "Test Author",
		PublishDate: time.Now(),
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func GetBookById(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	book := models.Book{
		Id:          uint(id),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Test Title",
		Author:      "Test Author",
		PublishDate: time.Now(),
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   book,
	})
}

func CreateBook(c echo.Context) error {
	var jsonBody models.Book
	err := c.Bind(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(jsonBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// created
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "created",
		"data":   jsonBody,
	})
}

func UpdateBook(c echo.Context) error {
	var jsonBody models.Book
	err := c.Bind(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jsonBody.Id = uint(id)

	if err := c.Validate(jsonBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	// updated
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "updated",
		"data":   jsonBody,
	})
}

func DeleteBook(c echo.Context) error {
	strId := c.Param("id")
	_, err := strconv.Atoi(strId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// deleted
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
	})
}
