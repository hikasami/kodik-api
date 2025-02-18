package models

// QualityEntry описывает отдельную запись качества видео, приходящую в ответе от API.
type QualityEntry struct {
	Title string `json:"title"` // Название качества (например, "WEB-DLRip 720p")
	Count int    `json:"count"` // Число материалов, имеющих данное качество
}

// QualitiesResponse представляет структуру ответа для эндпоинта /qualities/v2.
type QualitiesResponse struct {
	Time    string         `json:"time"`    // Время выполнения запроса (например, "5ms")
	Total   int            `json:"total"`   // Общее количество записей
	Results []QualityEntry `json:"results"` // Список записей с качествами видео
}
