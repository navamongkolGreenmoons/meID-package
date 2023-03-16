package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/navamongkolgreenmoons/meID-package/constants"
)

func (h *Http) Get(url string, headers *map[string]string, respBody interface{}) (int, error) {
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return constants.ZERO, reqErr
	}
	h.addHeaders(headers, req)
	resp, err := h.sendRequest(req, respBody)
	if err != nil {
		return http.StatusConflict, err
	}

	return resp.StatusCode, nil
}

func (h *Http) Post(url string, headers *map[string]string, reqBody interface{}, respBody interface{}) (int, error) {
	jsonValue, _ := json.Marshal(reqBody)
	req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if reqErr != nil {
		return constants.ZERO, reqErr
	}
	h.addHeaders(headers, req)
	resp, err := h.sendRequest(req, respBody)
	if err != nil {
		return http.StatusConflict, err
	}
	return resp.StatusCode, nil
}

func (h *Http) addHeaders(headers *map[string]string, request *http.Request) {
	if headers != nil {
		for key, value := range *headers {
			request.Header.Set(key, value)
		}
	}
}

func (h *Http) mapBytesToStruct(body io.ReadCloser, structBody interface{}) error {
	byteBody, err := io.ReadAll(body)
	fmt.Println(string(byteBody))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteBody, &structBody); err != nil {
		return err
	}
	return nil
}

func (h *Http) sendRequest(request *http.Request, response interface{}) (*http.Response, error) {
	client := new(http.Client)
	resp, respErr := client.Do(request)
	if respErr != nil {
		return nil, respErr
	}

	defer resp.Body.Close()
	if err := h.mapBytesToStruct(resp.Body, response); err != nil {
		return nil, err
	}
	return resp, nil
}
