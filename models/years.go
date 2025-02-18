package models

// YearEntry описывает отдельную запись года в ответе API.
type YearEntry struct {
	Year  int `json:"year"`  // Год материала
	Count int `json:"count"` // Количество материалов за данный год
}

// YearsResponse представляет структуру ответа от эндпоинта /years API Kodik.
type YearsResponse struct {
	Time    string      `json:"time"`    // Время выполнения запроса (например, "5ms")
	Total   int         `json:"total"`   // Общее количество лет
	Results []YearEntry `json:"results"` // Список записей с годом и количеством материалов
}
