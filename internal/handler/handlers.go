package handler

import (
	"encoding/json"

	models "github.com/Hirogava/ServiceBuyer/internal/model/request"
	dbErrors "github.com/Hirogava/ServiceBuyer/internal/errors/db"
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
	var req models.CountingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Logger.Error("Error decoding request body", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	data, err := manager.CountingServiceRequest(&req)
	if err != nil {
		if err == dbErrors.ErrNoRecordsFound {
			http.Error(w, "No records found", http.StatusNotFound)
			return
		} else {
			log.Logger.Error("Error counting service requests", err)
			http.Error(w, "Error counting service requests", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}
