package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Translations выполняет запрос к эндпоинту /translations/v2 API Kodik,
// используя параметры, заданные в структуре models.TranslationsParams.
func Translations(tp *models.TranslationsParams) (*models.TranslationsResponse, error) {
	var response models.TranslationsResponse
	params := tp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/translations/v2", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
