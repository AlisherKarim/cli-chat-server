package routes

import (
	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/go-chi/chi"
)

func NewAuthRouter(main_handler *handlers.MainHandler) *chi.Mux {
	authRoutes := chi.NewRouter()

	authRoutes.Post("/login", main_handler.Login)
	authRoutes.Post("/signup", main_handler.SignUp)

	return authRoutes
}