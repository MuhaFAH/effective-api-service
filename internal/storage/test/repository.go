package test

import (
	"context"
	"github.com/MuhaFAH/effective-api-service/internal/storage"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"sync"
	"time"
)

var (
	mu         sync.RWMutex
	primaryKey uint
)

func generateID() uint {
	mu.Lock()
	defer mu.Unlock()

	primaryKey++

	return primaryKey
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUser(ctx context.Context, user e.User) (e.User, error) {
	now := time.Now()

	user.ID = generateID()
	user.CreatedAt = now
	user.UpdatedAt = now

	return user, nil
}

func (r *Repository) ReadUser(ctx context.Context, id uint) (e.User, error) {
	now := time.Now()
	name, surname, patronymic, age, gender, national := "Ivan", "Ivanov", "Ivanovich", 15, "Male", "ru"

	user := e.User{
		ID: id, CreatedAt: now, UpdatedAt: now,
		Name: &name, Surname: &surname, Patronymic: &patronymic, Age: &age, Gender: &gender, Nationality: &national,
	}

	return user, nil
}

func (r *Repository) ReadUsersByFilter(ctx context.Context, filter storage.Filter) ([]e.User, error) {
	now := time.Now()

	name, surname, patronymic, age, nat, gender := "Alice", "Ivanova", "Ivanovna", 18, "UA", "female"
	user1 := e.User{
		ID: 1, CreatedAt: now, UpdatedAt: now,
		Name: &name, Surname: &surname, Patronymic: &patronymic, Age: &age, Gender: &gender, Nationality: &nat}

	name, surname, patronymic, age, nat, gender = "Artem", "Petrov", "Petrovich", 35, "RU", "male"
	user2 := e.User{
		ID: 2, CreatedAt: now, UpdatedAt: now,
		Name: &name, Surname: &surname, Patronymic: &patronymic, Age: &age, Gender: &gender, Nationality: &nat}

	return []e.User{user1, user2}, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id uint, user e.User) (e.User, error) {
	now := time.Now()

	user.ID = id
	user.UpdatedAt = now

	return user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id uint) (e.User, error) {
	now := time.Now()
	name, surname, patronymic, age, gender, national := "Alice", "Vasil", "Ivanovich", 18, "Female", "ua"

	user := e.User{
		ID: id, CreatedAt: now, UpdatedAt: now,
		Name: &name, Surname: &surname, Patronymic: &patronymic, Age: &age, Gender: &gender, Nationality: &national,
	}

	return user, nil
}
