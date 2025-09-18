package http

import (
handlers "github.com/Hirogava/ServiceBuyer/internal/handler"
db "github.com/Hirogava/ServiceBuyer/internal/repository/postgres"

"github.com/gorilla/mux"
httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(manager *db.Manager) *mux.Router {
r := mux.NewRouter()

// Swagger UI
r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

handlers.InitChatRoutes(r, manager)

return r
}
