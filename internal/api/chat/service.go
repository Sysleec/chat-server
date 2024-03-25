package chat

import (
	"sync"

	"github.com/Sysleec/chat-server/internal/service"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

type Chat struct {
	streams map[string]desc.ChatV1_ConnectChatServer
	m       sync.RWMutex
}

// Server is the chat server
type Server struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService

	chats  map[string]*Chat
	mxChat sync.RWMutex

	channels  map[string]chan *desc.Message
	mxChannel sync.RWMutex
}

// NewServer returns a new chat server
func NewServer(chatService service.ChatService) *Server {
	return &Server{
		chatService: chatService,
		chats:       make(map[string]*Chat),
		channels:    make(map[string]chan *desc.Message),
	}
}
