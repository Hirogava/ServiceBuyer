package http

import (
	handlers "github.com/Hirogava/ServiceBuyer/internal/handler"
	db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"

	"github.com/gorilla/mux"
)

func NewRouter(manager *db.Manager) *mux.Router {
	r := mux.NewRouter()

	handlers.InitChatRoutes(r, manager)

	return r
}