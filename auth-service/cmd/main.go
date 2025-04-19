package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"auth-service/internals/config"
	"auth-service/internals/handler"
	"auth-service/internals/repository/postgres"
	"auth-service/internals/service"

)

func main() {
	db, err := config.InitPostgres()
	if err != nil {
		log.Fatalf("error connecting to DB : %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()
	authHandler.SetupRoutes(router)

	port := "8081"
	log.Printf("auth service running on port: %s", port)

}
