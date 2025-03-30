package service

import (
	"context"
	"errors"
	"github.com/MuhaFAH/effective-api-service/internal/agent"
	"github.com/MuhaFAH/effective-api-service/internal/storage"
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
	enrichedUser, err := s.agent.EnrichInformation(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.storage.CreateUser(ctx, *enrichedUser)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil

}

func (s *Service) ReadUser(ctx context.Context, id uint) (*e.User, error) {
	user, err := s.storage.ReadUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) ReadUsersByFilter(ctx context.Context, filter storage.Filter) ([]e.User, error) {
	users, err := s.storage.ReadUsersByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) UpdateUser(ctx context.Context, id uint, user e.User) (*e.User, error) {
	user, err := s.storage.UpdateUser(ctx, id, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) DeleteUser(ctx context.Context, id uint) (*e.User, error) {
	deletedUser, err := s.storage.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &deletedUser, nil
}
