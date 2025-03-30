package gorm

import (
	"context"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

//TODO реализация методов репозитория горм

func (r *Repository) CreateUser(ctx context.Context, user e.User) (e.User, error) {

	// создание пользователя

	return e.User{}, nil
}

func (r *Repository) ReadUser(ctx context.Context, id uint) (e.User, error) {

	// чтение пользователя

	return e.User{}, nil
}

func (r *Repository) UpdateUser(ctx context.Context, user e.User) (e.User, error) {

	// обновление данных

	return e.User{}, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id uint) (e.User, error) {
	return e.User{}, nil
}
