package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/Sysleec/chat-server/internal/config"
	"github.com/Sysleec/chat-server/internal/config/env"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func (s *server) CreateChat(ctx context.Context, in *emptypb.Empty) (*desc.CreateChatResponse, error) {
	// Create chat with default values
	query, args, err := sq.Insert("chats").Columns("id", "created_at").
		Values(sq.Expr("DEFAULT"), sq.Expr("DEFAULT")).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to inser query: %v", err)
	}

	var id int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}
	log.Printf("inserted chat with id: %v", id)

	return &desc.CreateChatResponse{ChatId: &desc.Chat{ChatId: id}}, nil
}

func (s *server) GetChats(ctx context.Context, in *emptypb.Empty) (*desc.GetChatsResponse, error) {
	query, args, err := sq.Select("id", "created_at").From("chats").ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to select query: %v", err)
	}

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to select chats: %v", err)
	}
	defer rows.Close()

	var chats []*desc.ChatWithTime
	for rows.Next() {
		chat := &desc.ChatWithTime{}
		chat.ChatId = &desc.Chat{}
		var timestamp time.Time
		err := rows.Scan(&chat.ChatId.ChatId, &timestamp)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan chat: %v", err)
		}
		chat.Timestamp = timestamppb.New(timestamp)
		chats = append(chats, chat)
	}

	return &desc.GetChatsResponse{Chats: chats}, nil
}

func main() {
	flag.Parse()

	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to load pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to pg: %v", err)
	}
	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
