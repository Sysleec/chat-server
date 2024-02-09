package chat

import (
	"context"

	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateChat(ctx context.Context, req *emptypb.Empty) (*desc.CreateChatResponse, error) {
	id, err := s.chatService.CreateChat(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{
		Chat: &desc.Chat{ChatId: id},
	}, nil
}
