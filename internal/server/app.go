package server

import (
	"context"
	"github.com/MuhaFAH/effective-api-service/internal/agent"
	"github.com/MuhaFAH/effective-api-service/internal/config"
	"github.com/MuhaFAH/effective-api-service/internal/service"
	"github.com/MuhaFAH/effective-api-service/internal/storage/gorm"
)

//TODO добавить везде контекст

func RunApp(ctx context.Context, cfg config.Config) error {
	client, repo := agent.NewRestyAgent(), gorm.NewRepository(cfg.PostgresConfig)

	srv := service.NewService(client, repo)
	server := NewServer(ctx, srv)

	server.InitHandlers()

	err := server.Run(cfg.Address)
	if err != nil {
		return err
	}

	return nil
}
