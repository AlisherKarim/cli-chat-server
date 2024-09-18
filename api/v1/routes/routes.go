package routes

import (
	"net/http"

	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/go-chi/chi"
)

func RegisterRoutes(router *chi.Mux, storage db.Storage) {
	v1Router := chi.NewRouter()
	handler := handlers.NewHandler(storage)

	v1Router.Get("/healtz", handlers.HandlerReadiness)
	v1Router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			handler.GetUsers(w, r)
		} else {
			handler.GetUserById(w, r)
		}
	})
	v1Router.Post("/users", handler.CreateUser)

	RegisterAuthRoutes(v1Router)

	router.Mount("/api/v1", v1Router)
}

