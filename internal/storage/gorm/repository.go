package gorm

import (
	"context"
	"github.com/MuhaFAH/effective-api-service/internal/storage"
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

func (r *Repository) ReadUsersByFilter(ctx context.Context, filter storage.Filter) ([]e.User, error) {
	var users []e.User

	filters := map[string]interface{}{}
	if filter.Name != nil {
		filters["name"] = *filter.Name
	}
	if filter.Surname != nil {
		filters["surname"] = *filter.Surname
	}
	if filter.Patronymic != nil {
		filters["patronymic"] = *filter.Patronymic
	}
	if filter.Age != nil {
		filters["age"] = *filter.Age
	}
	if filter.Gender != nil {
		filters["gender"] = *filter.Gender
	}
	if filter.Nationality != nil {
		filters["nationality"] = *filter.Nationality
	}

	err := r.db.Scopes(paginate(filter)).Where(filters).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func paginate(filter storage.Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := *filter.Page
		if page <= 0 {
			page = 1
		}

		pageSize := *filter.Count
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
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
