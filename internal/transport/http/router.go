package http

import (
	handlers "github.com/Hirogava/ServiceBuyer/internal/handler"

	"github.com/gorilla/mux"
)

func NewChatRouter() *mux.Router {
	r := mux.NewRouter()

	handlers.InitChatRoutes(r)

	return r
}