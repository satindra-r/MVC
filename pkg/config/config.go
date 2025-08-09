package config

import (
	"fmt"
	"mvc/pkg/utils"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbHost     string
	dbUser     string
	dbPassword string
	database   string
	dbPort     string
	ServerPort string
	JWTSecret  string
}

var EnvConfig Config

func LoadEnvs() {
	var err = godotenv.Load()
	utils.PanicIfErr(err, "Error loading .env file")
	EnvConfig.ServerPort = os.Getenv("SERVER_PORT")
	EnvConfig.JWTSecret = os.Getenv("JWT_SECRET")
	EnvConfig.dbHost = os.Getenv("MYSQL_HOST")
	EnvConfig.dbUser = os.Getenv("MYSQL_USER")
	EnvConfig.dbPassword = os.Getenv("MYSQL_PASSWORD")
	EnvConfig.database = os.Getenv("MYSQL_DATABASE")
	EnvConfig.dbPort = os.Getenv("MYSQL_PORT")
}

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		EnvConfig.dbUser, EnvConfig.dbPassword, EnvConfig.dbHost, EnvConfig.dbPort, EnvConfig.database)
}
