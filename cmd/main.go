package main

import (
	"log"

	env "github.com/Hirogava/ServiceBuyer/internal/config/environment"
)

func main() {
	if err := env.LoadEnvFile("./.env"); err != nil {
		log.Fatalf("Ошибка загрузки env: %v", err)
	}
}
