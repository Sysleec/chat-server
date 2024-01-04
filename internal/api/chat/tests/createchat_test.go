package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Sysleec/chat-server/internal/api/chat"

	"github.com/Sysleec/chat-server/internal/service"
	serviceMocks "github.com/Sysleec/chat-server/internal/service/mocks"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *emptypb.Empty
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		serviceErr = fmt.Errorf("service error")

		req = &emptypb.Empty{}

		res = &desc.CreateChatResponse{
			Chat: &desc.Chat{
				ChatId: id,
			},
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateChatResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := chat.NewServer(userServiceMock)

			newID, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
