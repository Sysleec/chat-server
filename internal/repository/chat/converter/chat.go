package converter

import (
	"github.com/Sysleec/chat-server/internal/model"
	modelRepo "github.com/Sysleec/chat-server/internal/repository/chat/model"
)

func ToChatsFromRepo(repo []modelRepo.Chat) []model.Chat {
	chats := make([]model.Chat, len(repo))
	for i, chat := range repo {
		chats[i] = model.Chat{
			ID: chat.ID,
		}
	}
	return chats

}

func ToChatFromRepo(repo modelRepo.Chat) *model.Chat {
	return &model.Chat{
		ID: repo.ID,
	}
}
