package api

import (
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

// List выполняет запрос к эндпоинту /list API Kodik, используя параметры, заданные в структуре models.ListParams.
func List(lp *models.ListParams) (*models.ListResponse, error) {
	var response models.ListResponse
	params := lp.ToMap()
	err := client.DefaultClient.DoRequest("GET", "/list", params, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
