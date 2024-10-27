package routes

import (
	"github.com/alisherkarim/cli-chat-server/api/v1/handlers"
	"github.com/go-chi/chi"
)


func NewChatRouter(handler *handlers.MainHandler) *chi.Mux {
	chatRouter := chi.NewRouter();

	chatRouter.Post("/", handler.HandleCreate)
	chatRouter.Get("/", handler.HandleGetRooms)
	// chatRouter.Get("/{id}/messages", handler.HandleJoin)
	// chatRouter.HandleFunc("/{id}/join", handler.HandleJoin) // join the chat
	chatRouter.HandleFunc("/{id}/ws", handler.HandleWebSocketConnection) // join the chat
	// chatRouter.Get("/{id}", handler.HandleJoin)
	// chatRouter.HandleFunc("/{id}/send", handler.HandleJoin)

	return chatRouter
}