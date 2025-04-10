package storage

type Filter struct {
	Page  *int `json:"page"`
	Count *int `json:"count_in_page"`

	ID          *uint   `json:"id"`
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         *int    `json:"age"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
}
