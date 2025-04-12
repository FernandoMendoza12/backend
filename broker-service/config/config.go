package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	MongoUri  string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error while loading .env file for system variables")
	}

	return &Config{
		Port:      getEnv("PORT", "8080"),
		MongoUri:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret: getEnv("JWT_SECRET", "supersecreto"),
	}

}

func getEnv(key, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}
