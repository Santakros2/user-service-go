package main

import (
	"net/http"
	"users-service/internal/config"
	"users-service/internal/db"
	"users-service/internal/handlers"
	"users-service/internal/logger"
	"users-service/internal/middleware"
	"users-service/internal/repository"
	"users-service/internal/services"
)

func main() {
	logger.Init()
	cfg := config.LoadConfigMySQL()
	mux := http.NewServeMux()

	// oracleDB, err := db.ConnectOracle(cfg)
	mySqlDB, err := db.ConnectMySql(cfg)
	if err != nil {
		logger.Logger.Fatal("cannot connect to mysql:", err)
	}
	defer mySqlDB.Close()

	userRepo := repository.NewUserRepository(mySqlDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserById)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUserById)
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	logger.Logger.Println("users service running on port", cfg.AppPort)

	handler := middleware.Logging(mux)

	if err := http.ListenAndServe(":"+cfg.AppPort, handler); err != nil {
		logger.Logger.Fatal("server failed:", err)
	}
}
