package main

import "fmt"

var util *Util

func main() {

	// setupJwt := SetupJwt("", "")
	util = Init(nil)

	b := util.General.RandSaltKey()
	fmt.Println(b)

	url := "https://greenmoons-core-dev.azurewebsites.net/api/v1/token"
	type ReqBody struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}

	type RespBody struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
	}

	reqBody := ReqBody{
		PhoneNumber: "08123456789",
		Password:    "123456",
	}
	respBody := new(RespBody)
	result, err := util.Http.Post(url, nil, reqBody, respBody)
	fmt.Println(respBody, result, err)
}
