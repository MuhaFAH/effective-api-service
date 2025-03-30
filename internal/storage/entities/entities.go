package entities

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         *int    `json:"age"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
}
