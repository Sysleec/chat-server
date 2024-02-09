package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/model"
)

func (s *serv) GetChat(ctx context.Context, id int64) (*model.Chat, error) {
	chat, err := s.chatRepo.GetChat(ctx, id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
