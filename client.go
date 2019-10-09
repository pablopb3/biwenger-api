package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/magiconair/properties"
)

type Client struct {
	apiURL string
}

type request struct {
	endpoint string
	method   string
	body     []byte
	token    string
}

func (c Client) doRequest(r request) ([]byte, error) {
	req, err := http.NewRequest(r.method, c.apiURL+r.endpoint, bytes.NewBuffer(r.body))
	if err != nil {
		log.Println("error creating new request", err)
		return nil, err
	}

	// TODO improve this check
	if r.endpoint != loginURL {
		// TODO move it to init()
		p := properties.MustLoadFile("application.properties", properties.UTF8)
		req.Header.Set("authorization", "Bearer "+r.token)
		req.Header.Set("x-lang", "en")
		req.Header.Set("x-league", p.GetString("leagueId", ""))
		req.Header.Set("x-user", p.GetString("userId", ""))
		req.Header.Set("x-version", p.GetString("biwengerVersion", ""))
	}
	req.Header.Set("Content-Type", "application/json")

	Client := &http.Client{}
	resp, err := Client.Do(req)
	if err != nil {
		log.Println("error executing request", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading body from response", err)
		return nil, err
	}

	return body, nil
}
