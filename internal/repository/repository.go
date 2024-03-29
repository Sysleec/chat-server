package repository

import (
	"context"

	"github.com/Sysleec/chat-server/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserRepository is an interface for chat repository
type ChatRepository interface {
	CreateChat(ctx context.Context, in *emptypb.Empty) (int64, error)
	GetChat(ctx context.Context, id int64) (*model.Chat, error)
	GetChats(ctx context.Context, in *emptypb.Empty) ([]model.Chat, error)
	DeleteChat(ctx context.Context, cht *model.Chat) (*emptypb.Empty, error)
}
