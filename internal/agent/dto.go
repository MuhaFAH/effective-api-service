package agent

// ageResponse используется для получения данных о возрасте
type ageResponse struct {
	Count int    `json:"count"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
}

// genderResponse используется для получения данных о поле (гендере)
type genderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

// nationalizeResponse используется для получения данных о национальности
type nationalizeResponse struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []country `json:"country"`
}

// country используется для массива стран с вероятностью совпадения
type country struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
