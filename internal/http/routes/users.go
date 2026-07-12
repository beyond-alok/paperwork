package routes

import (
	"github.com/beyond-alok/paperwork/internal/http/handlers/users"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(userHandler *users.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", userHandler.GetUsers)
	r.Get("/{id}",userHandler.GetUserById)
	r.Get("/{id}",userHandler.DeleteUser)

	return  r
}
