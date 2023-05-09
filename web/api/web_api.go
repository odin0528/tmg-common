package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func sendRequest(method HTTP_METHOD, url string, header map[string]string, body interface{}, timeout time.Duration) ([]byte, error) {
	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return []byte{}, err
	}

	bodyReader := bytes.NewReader(bodyBytes)

	request, err := http.NewRequest(string(method), url, bodyReader)
	if err != nil {
		return []byte{}, err
	}

	for k, v := range header {
		request.Header.Set(k, v)
	}

	tr := http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Timeout:   timeout,
		Transport: &tr,
	}
	resp, err := client.Do(request)

	if err != nil {
		return []byte{}, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return respBytes, nil
}

func SendWebAPIAsync(method HTTP_METHOD, url string, header map[string]string, body interface{}, callback WebAPICallback, callerInfo interface{}) {
	go func(callback func([]byte, interface{}, error), callerInfo interface{}) {
		respBytes, err := sendRequest(method, url, header, body, (DEFAULT_API_TIME_OUT * time.Second))
		callback(respBytes, callerInfo, err)
	}(callback, callerInfo)
}

func SendWebAPI(method HTTP_METHOD, url string, header map[string]string, body interface{}) ([]byte, error) {
	return sendRequest(method, url, header, body, (DEFAULT_API_TIME_OUT * time.Second))
}
