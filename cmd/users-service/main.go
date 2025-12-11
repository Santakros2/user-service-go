package main

import (
	"log"
	"net/http"

	"users-service/internal/config"
	"users-service/internal/db"
	"users-service/internal/handlers"
	"users-service/internal/repository"
	"users-service/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	oracleDB, err := db.ConnectOracle(cfg)
	if err != nil {
		log.Fatal("Cannot connect to Oracle:", err)
	}

	userRepo := repository.NewUserRepository(oracleDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.CreateUser)

	log.Println("Users Service running on port", cfg.AppPort)
	http.ListenAndServe(":"+cfg.AppPort, nil)
}
