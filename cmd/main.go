package main

import (
	"os"

	"github.com/sirupsen/logrus"

	env "github.com/Hirogava/ServiceBuyer/internal/config/environment"
	log "github.com/Hirogava/ServiceBuyer/internal/config/logger"
	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"
	router "github.com/Hirogava/ServiceBuyer/internal/transport/http"
)

func main() {
	if err := env.LoadEnvFile("./.env"); err != nil {
		logrus.Fatalf("Ошибка загрузки env: %v", err)
	}

	log.LogInit()

	manager := db.NewManager("postgres", os.Getenv("POSTGRES_CONNECTION_STRING"))
	logrus.Info("postgres connected")
	db.Migrate(manager.Conn)
	logrus.Info("migrations completed")

	r := router.NewRouter(manager)

	s := router.NewServer(os.Getenv("SERVICE_SERVER_PORT"), r)
	if err := s.ListenAndServe(); err != nil {
		logrus.Fatalf("server startup error: %v", err)
	}
}
