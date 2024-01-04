package chat

import (
	"context"
	"fmt"

	"github.com/Sysleec/chat-server/internal/converter"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetChats(ctx context.Context, req *emptypb.Empty) (*desc.GetChatsResponse, error) {
	chats, err := s.chatService.GetChats(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("got chats: %+v\n", chats)

	return converter.ToChatsFromService(chats), nil
}
