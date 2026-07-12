package routes

import (
	"github.com/beyond-alok/paperwork/internal/http/handlers/auth"
	"github.com/go-chi/chi/v5"
)


func AuthRoutes(authHandler *auth.AuthHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/register", authHandler.Register)
	r.Post("/login",authHandler.Login)
	r.Get("/logout",authHandler.Logout)

	return r
}
 
