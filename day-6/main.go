package main

import (
	"day-6/database"
	"day-6/internal/middlewares"
	"day-6/internal/repositories"
	"day-6/internal/routes"
	"day-6/internal/services"
	"day-6/pkg/utils"
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
