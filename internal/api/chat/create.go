package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/converter"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.chatService.Create(ctx, converter.ToChatFromDescCreate(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
