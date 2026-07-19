package routes

import (
	"github.com/beyond-alok/paperwork/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

func TenantRoutes(authHandler *handlers.TenantHandler) chi.Router {
	r := chi.NewRouter()

	// TODO : implement auth middleware
	r.Post("/register", authHandler.Register)

	return r
}
