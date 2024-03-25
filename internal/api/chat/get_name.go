package chat

import (
	"context"
	"fmt"

	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) GetName(ctx context.Context, empty *empty.Empty) (*desc.GetNameResponse, error) {
	res, err := s.chatService.GetName(ctx)
	if err != nil {
		return nil, fmt.Errorf("cant get name %v", err.Error())
	}

	return &desc.GetNameResponse{
		Name: res,
	}, nil
}
