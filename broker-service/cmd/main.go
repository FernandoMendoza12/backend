package main

import (
	"fmt"
	"log"

	"broker-service/config"
	"broker-service/internals/adapters/api/server"
)

func main() {

	cfg := config.LoadConfig()
	fmt.Println("Initalizing server on port: ", cfg.Port)

	if err := server.StarServer(cfg); err != nil {
		log.Fatalf("Error while init server %v: ", err)
	}
}
