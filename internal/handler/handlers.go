package handler

import (
	"encoding/json"

	models "github.com/Hirogava/ServiceBuyer/internal/model/request"
	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"
	log "github.com/Hirogava/ServiceBuyer/internal/config/logger"

	"net/http"
)

func RecordHandler(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	var req models.ServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("Error decoding request body", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := manager.CreateServiceRequest(&req); err != nil {
		log.Logger.Error("Error creating service request", err)
		http.Error(w, "Error creating service request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

func CountingHandler(w http.ResponseWriter, r *http.Request, manager *db.Manager) {

}