package models

// ListResponse описывает ответ от эндпоинта /list API Kodik.
type ListResponse struct {
	Time     string     `json:"time"`
	Total    int        `json:"total"`
	PrevPage *string    `json:"prev_page"` // Может быть null
	NextPage *string    `json:"next_page"` // Может быть null
	Results  []Material `json:"results"`
}
