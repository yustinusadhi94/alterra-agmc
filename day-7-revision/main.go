package main

import (
	"day-7-revision/database"
	"day-7-revision/database/migration"
	"day-7-revision/internal/factory"
	"day-7-revision/internal/http"
	middlewares "day-7-revision/internal/middleware"
	utils "day-7-revision/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()

	migration.Migrate()

	f := factory.NewFactory()
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	middlewares.LogMiddleware(e)
	
	middlewares.InitJWTSecretKey()
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
