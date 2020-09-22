package core

import (
	"encoding/json"
)

type YamlStatusResponseData struct {
	CheckoutUrl string `json:"checkout_url"`
	Active      int    `json:"active"`
}

type YamlStatusResponse struct {
	Error   bool                   `json:"error`
	Data    YamlStatusResponseData `json:"data"`
	Message string                 `json:"message"`
}

func ActionStatus(email string, token string, filepath string) (yamlResponse YamlStatusResponse, url string, statusCode int, err error) {
	_, respBytes, url, statusCode, err := Upload(email, token, filepath, ACTION_STATUS)
	err = json.Unmarshal(respBytes, &yamlResponse)

	return yamlResponse, url, statusCode, err
}
