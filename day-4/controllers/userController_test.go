package controllers

import (
	"bytes"
	"day-4/config"
	"day-4/middlewares"
	"day-4/models"
	"day-4/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	err := godotenv.Load("../.mock.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.InitDB()

	config.DB.Save(&models.User{
		Model:    gorm.Model{ID: 1},
		Name:     "test name 1",
		Email:    "test_email@test.com",
		Password: "test_only",
	})

	config.DB.Save(&models.User{
		Model:    gorm.Model{ID: 2},
		Name:     "test name 2",
		Email:    "test_email_2@test.com",
		Password: "test_only",
	})

	middlewares.InitJWTSecretKey()
}

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")

	// assertions valid test case
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserById_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserById_invalid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("s")

	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateUser(t *testing.T) {
	userJSON := models.User{
		Name:     "new name",
		Email:    "new_emaillll@mail.com",
		Password: "new_password",
	}

	data, _ := json.Marshal(userJSON)
	reader := bytes.NewReader(data)
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")

	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

//func TestUpdateUser(t *testing.T) {
//	userJSON := models.User{
//		Name:     "new name",
//		Email:    "new_emaillll@mail.com",
//		Password: "new_password",
//	}
//
//	token := createJwtToken()
//	data, _ := json.Marshal(userJSON)
//	reader := bytes.NewReader(data)
//	e := echo.New()
//
//	e.Validator = &utils.CustomValidator{Validator: validator.New()}
//
//	req := httptest.NewRequest(http.MethodPut, "/", reader)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
//
//	jwt.Parse(token, )
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.Set("user", jwt.Token{Claims: })
//	c.SetPath("/v1/users/:id")
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	if assert.NoError(t, UpdateUser(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//	}
//}
//
//func TestDeleteUser(t *testing.T) {
//	e := echo.New()
//
//	req := httptest.NewRequest(http.MethodDelete, "/", nil)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/v1/users/:id")
//	c.SetParamNames("id")
//	c.SetParamValues("1")
//
//	if assert.NoError(t, DeleteUser(c)) {
//		assert.Equal(t, http.StatusOK, rec.Code)
//	}
//}
