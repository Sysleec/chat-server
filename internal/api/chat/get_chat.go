package chat

import (
	"context"
	"fmt"

	"github.com/Sysleec/chat-server/internal/converter"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

func (s *Server) GetChat(ctx context.Context, req *desc.GetChatRequest) (*desc.GetChatResponse, error) {
	chat, err := s.chatService.GetChat(ctx, req.GetChatId())
	if err != nil {
		return nil, err
	}

	fmt.Printf("got chat: %+v\n", chat)

	return converter.ToChatFromService(chat), nil
}
