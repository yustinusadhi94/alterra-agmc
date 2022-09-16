package controllers

import (
	"bytes"
	"day-4/models"
	"day-4/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetBooks_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	// assertions valid test case
	if assert.NoError(t, GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookById_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookById_invalid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("s")

	if assert.NoError(t, GetBookById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateBook(t *testing.T) {
	bookJSON := models.Book{
		Title:       "this is title",
		Author:      "this is author",
		PublishDate: time.Now(),
	}

	data, _ := json.Marshal(bookJSON)
	reader := bytes.NewReader(data)
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")

	if assert.NoError(t, CreateBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUpdateBook(t *testing.T) {
	bookJSON := models.Book{
		Title:       "this is title",
		Author:      "this is author",
		PublishDate: time.Now(),
	}

	data, _ := json.Marshal(bookJSON)
	reader := bytes.NewReader(data)
	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, UpdateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}