package app

import (
	"context"
	"log"

	"github.com/Sysleec/chat-server/internal/api/chat"
	"github.com/Sysleec/chat-server/internal/client/db"
	"github.com/Sysleec/chat-server/internal/client/db/pg"
	"github.com/Sysleec/chat-server/internal/client/db/transaction"
	"github.com/Sysleec/chat-server/internal/closer"
	"github.com/Sysleec/chat-server/internal/config"
	"github.com/Sysleec/chat-server/internal/config/env"
	"github.com/Sysleec/chat-server/internal/repository"
	chatRepository "github.com/Sysleec/chat-server/internal/repository/chat"
	"github.com/Sysleec/chat-server/internal/service"
	chatService "github.com/Sysleec/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatServ *chat.Server
}

// NewServiceProvider creates a new service provider
func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %s", err.Error())
		}

		closer.Add(client.Close)

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepo(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatServer(ctx context.Context) *chat.Server {
	if s.chatServ == nil {
		s.chatServ = chat.NewServer(s.ChatService(ctx))
	}

	return s.chatServ
}
