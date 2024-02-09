package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) CreateChat(ctx context.Context, req *emptypb.Empty) (int64, error) {
	var id int64

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepo.CreateChat(ctx, &emptypb.Empty{})
		if errTx != nil {
			return errTx
		}

		_, errTx = s.chatRepo.GetChats(ctx, &emptypb.Empty{})
		if errTx != nil {
			return errTx
		}

		return nil

	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
