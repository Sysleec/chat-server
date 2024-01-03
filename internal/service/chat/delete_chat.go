package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) DeleteChat(ctx context.Context, cht *model.Chat) (*emptypb.Empty, error) {
	_, err := s.chatRepo.DeleteChat(ctx, cht)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
