package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"
)

func InitChatRoutes(r *mux.Router, manager *db.Manager) {
	r.HandleFunc("/record", func(w http.ResponseWriter, r *http.Request) {
		RecordHandler(w, r, manager)
	}).Methods(http.MethodPost)

	r.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		CountingHandler(w, r, manager)
	}).Methods(http.MethodGet)
}
