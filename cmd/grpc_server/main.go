package main

import (
	"context"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/Sysleec/chat-server/internal/config"
	"github.com/Sysleec/chat-server/internal/config/env"
	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
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
	builderInsert := sq.Insert("chats").Suffix("RETURNING id")
	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var id int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}
	log.Printf("inserted chat with id: %v", id)

	return &desc.CreateChatResponse{ChatId: &desc.Chat{ChatId: id}}, nil
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
