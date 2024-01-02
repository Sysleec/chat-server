package repository

import (
	"context"

	"github.com/Sysleec/chat-server/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserRepository is an interface for chat repository
type ChatRepository interface {
	CreateChat(ctx context.Context, _ *emptypb.Empty) (int64, error)
	GetChats(ctx context.Context, _ *emptypb.Empty) ([]model.Chat, error)
}
