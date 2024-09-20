package routes

import (
	"net/http"

	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/go-chi/chi"
)

func NewRouter(storage db.Storage) *chi.Mux {
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

	v1Router.Mount("/auth", NewAuthRouter(handler))

	return v1Router
}

