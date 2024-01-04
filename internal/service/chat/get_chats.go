package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) GetChats(ctx context.Context, _ *emptypb.Empty) ([]model.Chat, error) {
	chats, err := s.chatRepo.GetChats(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return chats, nil
}
