package repository

import (
	"context"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, User e.User) (e.User, error)
	ReadUser(ctx context.Context, id uint) (e.User, error)
	UpdateUser(ctx context.Context, id uint, User e.User) (e.User, error)
	DeleteUser(ctx context.Context, id uint) (e.User, error)
}
