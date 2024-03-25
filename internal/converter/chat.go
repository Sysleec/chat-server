package converter

import (
	"github.com/Sysleec/chat-server/internal/model"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

func ToChatsFromService(chats []model.Chat) *desc.GetChatsResponse {
	resp := &desc.GetChatsResponse{}
	for _, chat := range chats {
		resp.Chats = append(resp.Chats, ToChat(chat))
	}
	return resp
}

func ToChat(chat model.Chat) *desc.Chat {
	return &desc.Chat{
		ChatId: chat.ID,
	}
}

func ToChatFromService(chat *model.Chat) *desc.GetChatResponse {
	return &desc.GetChatResponse{
		Chat: ToChat(*chat),
	}
}

func ToChatFromRepo(chat *desc.Chat) *model.Chat {
	return &model.Chat{
		ID: chat.ChatId,
	}
}

func ToChatFromDescCreate(chat *desc.CreateRequest) *model.Chat {
	return &model.Chat{
		ID:       0,
		Username: chat.GetUsername(),
		Password: chat.GetPassword(),
		Email:    chat.GetEmail(),
	}
}
