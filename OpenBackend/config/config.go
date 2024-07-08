package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DatabaseURL  string
	JWTSecret    string
	LogLevel     string
	LogFile      string
	LogToConsole bool
	Postgreshost string
	PostgresPort string
	PostgresUser string
	PostgresPassword string
	PostgresDB string
}

var AppConfig Config

func Init() {
	godotenv.Load()

	AppConfig = Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		LogLevel:     os.Getenv("LOG_LEVEL"),
		LogFile:      os.Getenv("LOG_FILE"),
		LogToConsole: os.Getenv("LOG_TO_CONSOLE") == "true",
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB: os.Getenv("POSTGRES_DB"),
		Postgreshost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}
}