package routes

import (
	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/go-chi/chi"
)

func RegisterAuthRoutes(r *chi.Mux) {
	authRoutes := chi.NewRouter()

	authRoutes.Post("/login", handlers.Login)
	authRoutes.Post("/signup", handlers.SignUp)

	r.Mount("/auth", authRoutes)
}