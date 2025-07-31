package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Env         string
	JWTSecret   string
	LogLevel    string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("WARNING: .env file not loaded:", err)
	}
}

// LoadConfig читает каждую переменную непосредственно из окружения,
// или использует дефолт, если переменная не установлена.
func LoadConfig() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("WARNING: DATABASE_URL not set, using empty string")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	jwt := os.Getenv("JWT_SECRET")
	if jwt == "" {
		fmt.Println("WARNING: JWT_SECRET not set")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return Config{
		DatabaseURL: dbURL,
		Port:        port,
		Env:         env,
		JWTSecret:   jwt,
		LogLevel:    logLevel,
	}
}
