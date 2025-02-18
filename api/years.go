package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Years выполняет запрос к эндпоинту /years API Kodik с параметрами models.YearsParams.
func Years(yp *models.YearsParams) (*models.YearsResponse, error) {
	var response models.YearsResponse
	params := yp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/years", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
