package models

// Country описывает отдельную страну в ответе API.
// Поле Title содержит название страны, а Count — количество материалов, связанных с ней.
type Country struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

// CountriesResponse представляет структуру ответа от эндпоинта /countries API Kodik.
type CountriesResponse struct {
	Time    string    `json:"time"`    // Время выполнения запроса (например, "5ms")
	Total   int       `json:"total"`   // Общее количество стран
	Results []Country `json:"results"` // Список стран с количеством материалов для каждой
}
