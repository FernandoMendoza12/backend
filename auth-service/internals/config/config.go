package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	JWTSecretKey string
}

func InitPostgres() (*gorm.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "authdb")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	var db *gorm.DB
	var err error

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database")
			return db, nil
		}
		wait := time.Duration(2<<i) * time.Second
		log.Printf("⚠️  Failed to connect to DB. Retrying in %v... (%d/%d)", wait, i+1, maxRetries)
		time.Sleep(wait)
	}

	return nil, err
}

var cfg *Config

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}

	_ = godotenv.Load()

	cfg = &Config{
		JWTSecretKey: getEnv("JWT_SECRET_KEY", "supersecretkey"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
