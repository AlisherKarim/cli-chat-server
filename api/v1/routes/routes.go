package routes

import (
	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/go-chi/chi"
)

func RegisterRoutes(router *chi.Mux) {
	v1Router := chi.NewRouter()

	v1Router.Get("/healtz", handlers.HandlerReadiness)
	v1Router.Get("/users", handlers.GetUsers)
	v1Router.Post("/users", handlers.CreateUser)

	RegisterAuthRoutes(v1Router)

	router.Mount("/api/v1", v1Router)
}

