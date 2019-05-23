package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doRequestAndGetJson(method string, url string, headers map[string]string, payload string) string {

	var jsonStr = []byte(payload)

	fmt.Printf("\n\nRequest: \n" + method + "\n" + url + "\n" + payload)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	responseBody := string(body)
	fmt.Printf("\nResponse: \n" + responseBody)
	return responseBody

}

func doRequestAndGetStruct(method string, url string, headers map[string]string, payload string, target interface{}) {

	jsonResponse := doRequestAndGetJson(method, url, headers, payload)
	json.Unmarshal([]byte(jsonResponse), &target)
}
