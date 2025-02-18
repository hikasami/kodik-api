package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Countries выполняет запрос к эндпоинту /countries API Kodik с параметрами models.CountriesParams.
func Countries(cp *models.CountriesParams) (*models.CountriesResponse, error) {
	var response models.CountriesResponse
	params := cp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/countries", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
