package handler

import (
	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"

	"net/http"
)

func RecordHandler(w http.ResponseWriter, r *http.Request, manager *db.Manager) {}

func CountingHandler(w http.ResponseWriter, r *http.Request, manager *db.Manager) {}