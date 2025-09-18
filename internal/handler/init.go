package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitChatRoutes(r *mux.Router) {
	r.HandleFunc("/record", func(w http.ResponseWriter, r *http.Request) {
		RecordHandler(w, r)
	}).Methods(http.MethodPost)

	r.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		CountingHandler(w, r)
	}).Methods(http.MethodGet)
}
