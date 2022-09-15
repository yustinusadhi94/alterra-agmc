package config

import (
	"day-3/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_HOST"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf(err.Error())
	}
}
