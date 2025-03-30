package gorm

import (
	"context"
	e "github.com/MuhaFAH/effective-api-service/internal/storage/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(config PostgresConfig) *Repository {
	db, err := Connect(config)
	if err != nil {
		panic(err)
	}

	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, user e.User) (e.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) ReadUser(ctx context.Context, id uint) (e.User, error) {
	var user e.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id uint, user e.User) (e.User, error) {
	r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"name":        gorm.Expr("COALESCE(?, users.name)", user.Name),
			"surname":     gorm.Expr("COALESCE(?, users.surname)", user.Surname),
			"patronymic":  gorm.Expr("COALESCE(?, users.patronymic)", user.Patronymic),
			"age":         gorm.Expr("COALESCE(?, users.age)", user.Age),
			"gender":      gorm.Expr("COALESCE(?, users.gender)", user.Gender),
			"nationality": gorm.Expr("COALESCE(?, users.nationality)", user.Nationality),
		})

	return user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, id uint) (e.User, error) {
	var user e.User
	err := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
