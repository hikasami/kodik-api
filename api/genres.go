package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// Genres выполняет запрос к эндпоинту /genres API Kodik с параметрами models.GenresParams.
func Genres(gp *models.GenresParams) (*models.GenresResponse, error) {
	var response models.GenresResponse
	params := gp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/genres", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
