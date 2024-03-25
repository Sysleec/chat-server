package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/auth"
)

func (s *serv) GetName(ctx context.Context) (string, error) {
	res, err := auth.GetName(ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}
