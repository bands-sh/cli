package core

import (
	"encoding/json"
)

type YamlUpResponseData struct {
	CheckoutUrl string `json:"checkout_url"`
}

type YamlUpResponse struct {
	Error   bool               `json:"error`
	Data    YamlUpResponseData `json:"data"`
	Message string             `json:"message"`
}

func ActionUp(email string, token string, filepath string) (yamlResponse YamlUpResponse, url string, statusCode int, err error) {
	_, respBytes, url, statusCode, err := Upload(email, token, filepath, ACTION_UP)
	err = json.Unmarshal(respBytes, &yamlResponse)

	return yamlResponse, url, statusCode, err
}
