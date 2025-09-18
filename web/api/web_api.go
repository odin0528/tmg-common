package api

import (
	"bytes"
	"fmt"
	"xxx/common/logs"
	"xxx/common/redis"

	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func InitUserAgentSwitch() {
	enableCustomUserAgent.Store(false) // 預設關閉

	// 初始化時先讀 Redis
	v, ok := redis.GetString(REDIS_KEY_ENABLE_CUSTOM_USER_AGENT)
	if ok && v == "1" {
		enableCustomUserAgent.Store(true)
	}

	// 定時刷新 Redis 開關
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			v, ok := redis.GetString(REDIS_KEY_ENABLE_CUSTOM_USER_AGENT)
			if ok {
				enableCustomUserAgent.Store(v == "1")
			}
		}
	}()

	logs.Info(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_API, fmt.Sprintf("InitUserAgentSwitch completed: %v", time.Now()), map[string]interface{}{})
}

func sendRequest(method HTTP_METHOD, url string, header map[string]string, body interface{}, timeout time.Duration) ([]byte, error) {
	var bodyBytes []byte
	var err error

	if body != nil {
		bodyBytes, err = json.Marshal(body)

		if err != nil {
			return []byte{}, err
		}
	}

	bodyReader := bytes.NewReader(bodyBytes)

	request, err := http.NewRequest(string(method), url, bodyReader)
	if err != nil {
		return []byte{}, err
	}

	for k, v := range header {
		request.Header.Set(k, v)
	}

	if enableCustomUserAgent.Load() {
		request.Header.Set(HEADER_KEY_USER_AGENT, CUSTOM_USER_AGENT)
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

func SendWebAPIWithSpecifySec(method HTTP_METHOD, url string, header map[string]string, body interface{}, timeout time.Duration) ([]byte, error) {
	return sendRequest(method, url, header, body, timeout)
}

func sendRequestGeneral(method HTTP_METHOD, url string, header map[string]string, body []byte, timeout time.Duration) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

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

	if resp == nil {
		return nil, errors.New("http response is nil")
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return respBytes, nil
}

func SendWebAPIGeneral(method HTTP_METHOD, url string, header map[string]string, body []byte) ([]byte, error) {
	return sendRequestGeneral(method, url, header, body, DEFAULT_API_TIME_OUT*time.Second)
}
