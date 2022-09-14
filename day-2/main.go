package main

import (
	"day-2/config"
	"day-2/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.InitDB()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
