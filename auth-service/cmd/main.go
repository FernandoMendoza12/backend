package main

import (
	"log"
)

func main() {
	db, err := config.InitPostgres()
	if err != nil {
		log.Fatalf("error connecting to DB : %v", err)
	}

	userRepo := postgres.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	router := handler.NewRouter(authService)

	port := "8081"
	log.Printf("auth service running on port: %s", port)
	router.Run(":" + port)
}
