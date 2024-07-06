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
	}
}