package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Client = &http.Client{Timeout: 30 * time.Second}

func Get(urlStr string) []byte {
	resp, _ := http.Get(urlStr)
	defer resp.Body.Close()
	str, _ := ioutil.ReadAll(resp.Body)
	return str
}

func Http(method string, urlStr string, paramBytes []byte) ([]byte, int) {
	return Http4(method, urlStr, paramBytes, "application/x-www-form-urlencoded")
}

func JsonPost(urlStr string, paramBytes []byte) ([]byte, int) {
	return Http4("POST", urlStr, paramBytes, "application/json")
}

func Http4(method string, urlStr string, paramBytes []byte, cType string) ([]byte, int) {
	headers := map[string]string{"Content-Type": cType}
	return HttpWithHeaders(method, urlStr, paramBytes, headers)
}

func HttpWithHeaders(method string, urlStr string, paramBytes []byte, headers map[string]string) ([]byte, int) {
	req, reqEr := http.NewRequest(method, urlStr, bytes.NewBuffer(paramBytes))
	if reqEr == nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}

		resp, err := Client.Do(req)

		if err != nil {
			return nil, 500
		} else {
			defer resp.Body.Close()
			str, _ := ioutil.ReadAll(resp.Body)

			return str, resp.StatusCode
		}
	} else {
		log.Println("Error! ", reqEr, " ", req)
	}

	return nil, 0
}
