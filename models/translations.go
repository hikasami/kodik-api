package models

// TranslationEntry описывает отдельную запись озвучки в ответе API.
type TranslationEntry struct {
	ID    int    `json:"id"`    // Уникальный идентификатор озвучки
	Title string `json:"title"` // Название озвучки
	Count int    `json:"count"` // Число материалов, в которых используется данная озвучка
}

// TranslationsResponse представляет структуру ответа эндпоинта /translations/v2.
type TranslationsResponse struct {
	Time    string             `json:"time"`    // Время выполнения запроса (например, "5ms")
	Total   int                `json:"total"`   // Общее количество записей
	Results []TranslationEntry `json:"results"` // Список записей (озвучек)
}
