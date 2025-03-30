package server

type GetUserExample struct {
	Name    *string `json:"name" example:"Ivan"`
	Surname *string `json:"surname" example:"Ivanov"`
}

type UpdateUserExample struct {
	Name *string `json:"name" example:"Alice"`
	Age  *int    `json:"age" example:"33"`
}

type GetUsersExample struct {
	Page  *int `json:"page" example:"1"`
	Count *int `json:"count_in_page" example:"3"`

	Gender *string `json:"gender" example:"male"`
}

type OkResponseExample struct {
	Status int `json:"status" example:"200"`
}

type ErrorResponseExample struct {
	Error  string `json:"error" example:"Bad Request"`
	Status int    `json:"status" example:"400"`
}
