package core

import (
	"bands/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	yamlLib "github.com/ghodss/yaml"
)

type Yaml struct {
	Yaml string `json:"yaml,omitempty"`
	Url  string `json:"checkout,omitempty"`
}

type AccountResponseData struct {
	ApiAccessToken     string `json:"access_key"`
	ApiActivationToken string `json:"activation_token"`
	Active             bool   `json:"active"`
}

type AccountResponse struct {
	Error bool                `json:"error`
	Data  AccountResponseData `json:"data"`
}

type ActionResponse struct {
	Error   bool   `json:"error`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type Account struct {
	Email  string `json:"email,omitempty"`
	Token  string `json:"token,omitempty"`
	Active bool   `json:"active,omitempty"`
}

var (
	apiURL        = "https://api.bands.sh"
	ACTION_UP     = "up"
	ACTION_DOWN   = "down"
	ACTION_STATUS = "status"
	respData      = &ActionResponse{}
)

func AccountCreate(email string, forced bool, debug bool) (accResp AccountResponse, statusCode int, err error) {
	var acc Account
	acc.Email = email
	params, err := json.Marshal(acc)
	url := apiURL + "/api/accounts/"

	if forced {
		url = url + "?force=true"
	}

	respBytes, statusCode := utils.JsonPost(url, params)

	if statusCode != 200 {
		return accResp, statusCode, errors.New(url)
	}

	if debug == true {
		fmt.Println("[debug]", url, "=>", string(respBytes))
	}

	err = json.Unmarshal(respBytes, &accResp)

	return accResp, statusCode, err
}

func Upload(email string, token string, filepath string, action string) (respData ActionResponse, respBytes []byte, url string, statusCode int, err error) {
	data, err := ioutil.ReadFile(filepath)
	params, err := yamlLib.YAMLToJSON(data)
	url = apiURL + "/api/action/" + action + "/?data_format=json&api_key=" + token
	respBytes, statusCode = utils.JsonPost(url, params)
	_ = json.Unmarshal([]byte(respBytes), &respData)

	return respData, respBytes, url, statusCode, err
}
