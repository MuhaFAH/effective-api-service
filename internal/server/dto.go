package server

import "github.com/MuhaFAH/effective-api-service/internal/storage"

type UsersRequest struct {
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

func convertIntoFilter(req UsersRequest) storage.Filter {
	if req.Page == nil {
		page := 1
		req.Page = &page
	}
	if req.Count == nil {
		count := 10
		req.Count = &count
	}

	return storage.Filter{
		Page:  req.Page,
		Count: req.Count,

		ID:          req.ID,
		Name:        req.Name,
		Surname:     req.Surname,
		Patronymic:  req.Patronymic,
		Age:         req.Age,
		Gender:      req.Gender,
		Nationality: req.Nationality,
	}
}
