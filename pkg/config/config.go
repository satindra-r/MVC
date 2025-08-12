package config

import (
	"fmt"
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
	_ = godotenv.Load()
	EnvConfig.ServerPort = "8090"
	EnvConfig.JWTSecret = os.Getenv("JWT_SECRET")
	EnvConfig.dbHost = os.Getenv("MYSQL_HOST")
	EnvConfig.dbUser = os.Getenv("MYSQL_USER")
	EnvConfig.dbPassword = os.Getenv("MYSQL_PASSWORD")
	EnvConfig.dbPort = os.Getenv("MYSQL_PORT")
}

func GetConnectionString(database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		EnvConfig.dbUser, EnvConfig.dbPassword, EnvConfig.dbHost, EnvConfig.dbPort, database)
}
