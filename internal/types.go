package types

import (
	"github.com/alisherkarim/cli-chat-server/models"
	"github.com/alisherkarim/cli-chat-server/ws"
)

type ChatRoom struct {
	DataBaseModel models.ChatRoom
	Hub *ws.Hub
}