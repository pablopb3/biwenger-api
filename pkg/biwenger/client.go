package biwenger

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/magiconair/properties"
)

const baseURL = "https://biwenger.as.com/api/v2"

type Client struct{}

type Request struct {
	Endpoint string
	Method   string
	Body     []byte
	Token    string
}

func (c Client) DoRequest(r Request) ([]byte, error) {
	req, err := http.NewRequest(r.Method, baseURL+r.Endpoint, bytes.NewBuffer(r.Body))
	if err != nil {
		log.Println("error creating new request", err)
		return nil, err
	}

	// TODO improve this check
	if r.Endpoint != "/auth/login" {
		// TODO move it to init()
		p := properties.MustLoadFile("application.properties", properties.UTF8)
		req.Header.Set("authorization", r.Token)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error not statusOK, StatusCode: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading body from response", err)
		return nil, err
	}

	return body, nil
}
