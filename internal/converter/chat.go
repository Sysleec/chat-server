package converter

import (
	"github.com/Sysleec/chat-server/internal/model"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
)

func ToChatsFromService(chats []model.Chat) *desc.GetChatsResponse {
	resp := &desc.GetChatsResponse{}
	for _, chat := range chats {
		resp.Chats = append(resp.Chats, ToChatFromService(chat))
	}
	return resp
}

func ToChatFromService(chat model.Chat) *desc.Chat {
	return &desc.Chat{
		ChatId: chat.ID,
	}
}
