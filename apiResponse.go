package main

import "encoding/json"

func SendApiResponse(response interface{}) string {
	return SendApiResponseWithMessage(response, "")
}

func SendApiResponseWithMessage(response interface{}, messages interface{}) string {
	resp := BiwengerApiResponse{"OK", 200, messages, response}
	jsonResp,_ := json.Marshal(resp)
	return string(jsonResp)
}

type BiwengerApiResponse struct {
	Status  string 	  		`json:"status"`
	Code 	int32			`json:"code"`
	Messages interface{}	`json:"messages"`
	Data  interface{}		`json:"data"`
}

