package main

import (
	"os"

	log "github.com/Hirogava/ServiceBuyer/internal/config/logger"
	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"
	router "github.com/Hirogava/ServiceBuyer/internal/transport/http"
	_ "github.com/Hirogava/ServiceBuyer/docs"
)

// @title ServiceBuyer API
// @version 1.0
// @description ServiceBuyer API - сервис для управления подписками пользователей
// @host localhost:8080
// @BasePath /
func main() {
	log.LogInit()

	postgresConnStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	serverPort := os.Getenv("SERVICE_SERVER_PORT")

	if postgresConnStr == "" {
		log.Logger.Fatalf("POSTGRES_CONNECTION_STRING environment variable is required")
	}
	if serverPort == "" {
		serverPort = "8080"
	}

	manager := db.NewManager("postgres", postgresConnStr)
	log.Logger.Info("postgres connected")
	db.Migrate(manager.Conn)
	log.Logger.Info("migrations completed")

	r := router.NewRouter(manager)

	s := router.NewServer(serverPort, r)
	log.Logger.Infof("Starting server on port %s", serverPort)
	if err := s.ListenAndServe(); err != nil {
		log.Logger.Fatalf("server startup error: %v", err)
	}
}
