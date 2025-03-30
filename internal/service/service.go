package service

import (
	"context"
	"errors"
	"github.com/MuhaFAH/effective-api-service/internal/agent"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"github.com/MuhaFAH/effective-api-service/internal/storage/repository"
)

type Service struct {
	agent   agent.Agent
	storage repository.Repository
}

func NewService(agent agent.Agent, storage repository.Repository) *Service {
	return &Service{agent: agent, storage: storage}
}

func (s *Service) CreateUser(ctx context.Context, user e.User) (*e.User, error) {
	if user.Name == nil || *user.Name == "" {
		return nil, errors.New("user name is required parameter")
	}
	if user.Surname == nil || *user.Surname == "" {
		return nil, errors.New("user surname is required parameter")
	}

	createdUser, err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	enrichedUser, err := s.agent.EnrichInformation(ctx, createdUser)
	if err != nil {
		return nil, err
	}

	return enrichedUser, nil

}

func (s *Service) ReadUser(ctx context.Context, id uint) (*e.User, error) {
	user, err := s.storage.ReadUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user e.User) (*e.User, error) {
	user, err := s.storage.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id uint) error {
	err := s.storage.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
