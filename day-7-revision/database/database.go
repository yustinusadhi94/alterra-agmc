package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
)

var (
	dbConn *gorm.DB
	// we use sync.Once to make sure we create connection only once
	once sync.Once
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}
)

func (m *mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Pass,
		m.Host,
		m.Port,
		m.Name,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func CreateConnection() *gorm.DB {
	conf := dbConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	mysql := mysqlConfig{dbConfig: conf}

	once.Do(func() {
		mysql.Connect()
		//postgres.Connect()
	})

	return dbConn
}

func InitMigrate(structs ...interface{}) {
	for _, value := range structs {
		err := dbConn.AutoMigrate(value)
		if err != nil {
			log.Println("error migration")
		}
	}
}

// GetConnection is a faction for return connection or return value dbConn
// because we set var dbConn is private
func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
