package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/converter"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	_, err := s.chatService.DeleteChat(ctx, converter.ToChatFromRepo(req.GetChat()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
