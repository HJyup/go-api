package config

import (
	"fmt"
	"github.com/lpernett/godotenv"
	"os"
	"strconv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
	JWTExpiration int64
	JWTSecret     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return Config{}
	}

	config := Config{
		PublicHost:    getEnv("PUBLIC_HOST", "http://localhost"),
		Port:          getEnv("PORT", "8080"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "password123"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:        getEnv("DB_NAME", "go_api"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600*24),
		JWTSecret:     getEnv("JWT_SECRET", "secret"),
	}

	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return intValue
	}

	return fallback
}
