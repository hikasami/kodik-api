package models

// Genre описывает отдельный жанр.
// Поле Title содержит название жанра, а Count — количество материалов, связанных с ним.
type Genre struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

// GenresResponse представляет структуру ответа от эндпоинта /genres API Kodik.
type GenresResponse struct {
	Time    string  `json:"time"`    // Время выполнения запроса (например, "5ms")
	Total   int     `json:"total"`   // Общее количество жанров
	Results []Genre `json:"results"` // Список жанров с количеством материалов для каждого
}
