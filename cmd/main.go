package main

import (
	"log/slog"
	"net/http"

	"github.com/beyond-alok/paperwork/internal/config"
	"github.com/beyond-alok/paperwork/internal/http/handlers"
	"github.com/beyond-alok/paperwork/internal/http/routes"
	"github.com/beyond-alok/paperwork/internal/service"
	"github.com/beyond-alok/paperwork/internal/storage/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := config.MustLoad()

	// setups database
	conn,err := postgresql.Connect(cfg.DBConfig)
	if err != nil {
		return 
	}
	slog.Info("database: postgresql connected successfully")

	// initialises each resources repo
	userRepo := postgresql.NewUserRepo(conn)

	authService := service.NewAuthService(userRepo)

	// initialises each resources handler
	authHandler := handlers.NewAuthHandler(authService)


	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Mount("/auth",routes.AuthRoutes(authHandler))

	http.ListenAndServe(cfg.HttpServerConfig.Port, r)
}


