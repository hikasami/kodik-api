package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Search выполняет запрос к эндпоинту /search API Kodik, используя параметры, заданные в структуре models.SearchParams.
func Search(sp *models.SearchParams) (*models.SearchResponse, error) {
	var response models.SearchResponse
	params := sp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/search", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
