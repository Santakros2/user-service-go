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
	cfg := config.LoadConfigMySQL()

	// oracleDB, err := db.ConnectOracle(cfg)
	mySqlDB, err := db.ConnectMySql(cfg)
	if err != nil {
		log.Fatal("Cannot connect to Oracle:", err)
	}
	defer mySqlDB.Close()

	userRepo := repository.NewUserRepository(mySqlDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.Users)

	log.Println("Users Service running on port", cfg.AppPort)
	http.ListenAndServe(":"+cfg.AppPort, nil)
	// fmt.Scanln()
}
