package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Qualities выполняет запрос к эндпоинту /qualities/v2 API Kodik, используя параметры models.QualitiesParams.
func Qualities(qp *models.QualitiesParams) (*models.QualitiesResponse, error) {
	var response models.QualitiesResponse
	params := qp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/qualities/v2", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
