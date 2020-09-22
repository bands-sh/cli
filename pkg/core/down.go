package core

import (
	"encoding/json"
)

type YamlDownResponseData struct {
	Done string `json:"done"`
}

type YamlDownResponse struct {
	Error   bool                 `json:"error`
	Data    YamlDownResponseData `json:"data"`
	Message string               `json:"message"`
}

func ActionDown(email string, token string, filepath string) (yamlResponse YamlDownResponse, url string, statusCode int, err error) {
	_, respBytes, url, statusCode, err := Upload(email, token, filepath, ACTION_DOWN)
	err = json.Unmarshal(respBytes, &yamlResponse)

	return yamlResponse, url, statusCode, err
}
