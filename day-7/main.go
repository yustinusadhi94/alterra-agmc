package main

import (
	"day-7/database"
	"day-7/internal/entities"
	"day-7/internal/middlewares"
	"day-7/internal/repositories"
	"day-7/internal/routes"
	"day-7/internal/services"
	"day-7/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	dbConn := database.CreateConnection()
	database.InitMigrate(entities.User{})
	// skip migration stuffs

	// initialize repositories
	userRepo := repositories.NewUserRepository(dbConn)

	// initialize services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)

	// initialize routes
	authRoute := routes.NewAuthRoute(authService)
	userRoute := routes.NewUserRoute(userService)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	middlewares.InitJWTSecretKey()

	// initialize http
	e = authRoute.AddAuthRoute(e)
	e = userRoute.AddUserRoutes(e, middlewares.ValidateToken())

	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

}
