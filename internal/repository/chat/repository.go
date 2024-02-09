package chat

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/Sysleec/chat-server/internal/client/db"
	"github.com/Sysleec/chat-server/internal/model"
	"github.com/Sysleec/chat-server/internal/repository"
	"github.com/Sysleec/chat-server/internal/repository/chat/converter"
	modelRepo "github.com/Sysleec/chat-server/internal/repository/chat/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	tableName = "chats"

	idColumn        = "id"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

func NewRepo(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) CreateChat(ctx context.Context, req *emptypb.Empty) (int64, error) {
	// Create chat with default values
	query, args, err := sq.Insert(tableName).Columns(idColumn, createdAtColumn).
		Values(sq.Expr("DEFAULT"), sq.Expr("DEFAULT")).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to inser query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Create_chat",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}
	log.Printf("inserted chat with id: %v", id)

	return id, nil
}

// GetChats returns all chats
func (r *repo) GetChats(ctx context.Context, req *emptypb.Empty) ([]model.Chat, error) {
	query, args, err := sq.Select(idColumn, createdAtColumn).From(tableName).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to select query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Get_chats",
		QueryRaw: query,
	}

	var chats []modelRepo.Chat
	err = r.db.DB().ScanAllContext(ctx, &chats, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get chats: %v", err)
	}

	return converter.ToChatsFromRepo(chats), nil
}

func (r *repo) GetChat(ctx context.Context, id int64) (*model.Chat, error) {
	query, args, err := sq.Select(idColumn, createdAtColumn).From(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumn: id}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to select query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Get_chat",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().ScanOneContext(ctx, &chat, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat: %v", err)
	}

	return converter.ToChatFromRepo(chat), nil
}

func (r *repo) DeleteChat(ctx context.Context, cht *model.Chat) (*emptypb.Empty, error) {
	query, args, err := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumn: cht.ID}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to delete query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.Delete_chat",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to delete chat: %v", err)
	}

	return &emptypb.Empty{}, nil
}
