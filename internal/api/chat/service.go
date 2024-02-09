package chat

import (
	"github.com/Sysleec/chat-server/internal/service"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

// Server is the chat server
type Server struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewServer returns a new chat server
func NewServer(chatService service.ChatService) *Server {
	return &Server{chatService: chatService}
}
