package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"mvc/pkg/utils"
	"os"
)

var dbHost string

var dbUser string

var dbPassword string

var database string

var dbPort string

var ServerPort string

var JWTSecret string

func LoadEnvs() {
	var err = godotenv.Load()
	utils.PanicIfErr(err, "Error loading .env file")
	if err == nil {
		ServerPort = os.Getenv("SERVER_PORT")
		JWTSecret = os.Getenv("JWT_SECRET")
		dbHost = os.Getenv("MYSQL_HOST")
		dbUser = os.Getenv("MYSQL_USER")
		dbPassword = os.Getenv("MYSQL_PASSWORD")
		database = os.Getenv("MYSQL_DATABASE")
		dbPort = os.Getenv("MYSQL_PORT")
	}
}

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, database)
}
