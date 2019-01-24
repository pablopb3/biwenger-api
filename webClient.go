package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func doRequestAndGetJson(method string, url string, headers map[string]string, payload string) string {

	var jsonStr = []byte(payload)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func doRequestAndGetStruct(method string, url string, headers map[string]string, payload string, target interface{}) {

	jsonResponse := doRequestAndGetJson(method, url, headers, payload)
	json.Unmarshal([]byte(jsonResponse), &target)
}
