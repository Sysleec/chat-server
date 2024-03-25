package chat

import (
	"context"
	"strconv"

	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	s.mxChannel.RLock()
	id := strconv.Itoa(int(req.GetChat().ChatId))
	chatChan, ok := s.channels[id]
	s.mxChannel.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "chat not found")
	}

	chatChan <- req.GetMessage()

	return &emptypb.Empty{}, nil
}
