package main

import (
	"day-4/config"
	"day-4/middlewares"
	"day-4/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load env files
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// init the database
	config.InitDB()

	// init jwt middleware (secret key)
	middlewares.InitJWTSecretKey()

	// create routes
	e := routes.New()

	// use logging middleware after routes
	middlewares.LogMiddleware(e)

	// start the server
	e.Logger.Fatal(e.Start(":8080"))
}
