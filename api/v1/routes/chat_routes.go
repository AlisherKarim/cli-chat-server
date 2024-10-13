package routes

import (
	"github.com/alisherkarim/cli-chat-server/pkg/websocket"
	"github.com/go-chi/chi"
)


func NewChatRouter() *chi.Mux {
	chatRouter := chi.NewRouter();

	chatRouter.HandleFunc("/connect", websocket.HandleConnections)
	// chatRouter.Post("/create")

	return chatRouter
}