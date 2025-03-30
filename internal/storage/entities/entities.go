package entities

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        *string `json:"name" example:"Ivan"`
	Surname     *string `json:"surname" example:"Ivanov"`
	Patronymic  *string `json:"patronymic" example:"Yovanovitch"`
	Age         *int    `json:"age" example:"18"`
	Gender      *string `json:"gender" example:"male"`
	Nationality *string `json:"nationality" example:"RU"`
}
