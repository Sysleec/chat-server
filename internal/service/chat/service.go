package chat

import (
	"github.com/Sysleec/chat-server/internal/client/db"
	"github.com/Sysleec/chat-server/internal/repository"
	def "github.com/Sysleec/chat-server/internal/service"
)

var _ def.ChatService = (*serv)(nil)

type serv struct {
	chatRepo  repository.ChatRepository
	txManager db.TxManager
}

// NewService returns a new user service
func NewService(chatRepo repository.ChatRepository,
	txManager db.TxManager) *serv {
	return &serv{
		chatRepo:  chatRepo,
		txManager: txManager,
	}
}
