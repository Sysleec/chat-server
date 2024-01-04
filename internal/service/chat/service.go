package chat

import (
	"github.com/Sysleec/chat-server/internal/client/db"
	"github.com/Sysleec/chat-server/internal/repository"
	"github.com/Sysleec/chat-server/internal/service"
)

var _ service.ChatService = (*serv)(nil)

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

func NewMockService(deps ...interface{}) service.ChatService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.ChatRepository:
			srv.chatRepo = s
		}
	}
	return &srv
}

func NewMockTxService(deps ...interface{}) *serv {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case serv:
			srv = s
		}
	}
	return &srv
}
