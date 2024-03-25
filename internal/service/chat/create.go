package chat

import (
	"context"

	"github.com/Sysleec/chat-server/internal/connect"
	"github.com/Sysleec/chat-server/internal/model"
	auth "github.com/Sysleec/chat-server/pkg/user_v1"
)

func (s *serv) Create(ctx context.Context, info *model.Chat) (int64, error) {
	//id, err := s.repo.Create(ctx, info)
	conn, err := connect.AuthServer()
	if err != nil {
		return 0, err
	}
	client := auth.NewUserV1Client(conn)
	defer conn.Close()
	resp, err := client.Create(ctx, &auth.CreateRequest{
		Name:            info.Username,
		Email:           info.Email,
		Password:        info.Password,
		PasswordConfirm: info.Password,
		Role:            auth.Role_user,
	})

	if err != nil {
		return 0, err
	}

	return resp.GetId(), nil
}
