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
	mux := http.NewServeMux()

	// oracleDB, err := db.ConnectOracle(cfg)
	mySqlDB, err := db.ConnectMySql(cfg)
	if err != nil {
		log.Fatal("Cannot connect to Oracle:", err)
	}
	defer mySqlDB.Close()

	userRepo := repository.NewUserRepository(mySqlDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /user/{id}", userHandler.GetUserById)
	mux.HandleFunc("PUT /user/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /user/{id}", userHandler.DeleteUserById)
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	log.Println("Users Service running on port", cfg.AppPort)

	if err := http.ListenAndServe(":"+cfg.AppPort, mux); err != nil {
		log.Fatal("Server failed:", err)
	}
}
