package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const homeURL = "/home"

type Home struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}

type Data struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type  string `json:"type"`
	Date  int    `json:"date"`
	Round Round  `json:"round"`
}

type Round struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Short string `json:"short"`
}

func (cli Client) getDaysToNextRound(c *gin.Context) {
	req := request{
		method:   "GET",
		endpoint: homeURL,
		token:    c.GetHeader("authorization"),
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for home", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var home Home
	err = json.Unmarshal(body, &home)
	if err != nil {
		log.Println("error unmarshalling response body to home", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	daysToNextRound, err := calculateDays(home.Data.Events)
	if err != nil {
		log.Println("error calculating days", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, daysToNextRound)
}

func calculateDays(events []Event) (int, error) {
	e := func(events []Event) *Event {
		for _, event := range events {
			if event.Type == "roundStart" { // TODO move to constant
				return &event
			}
		}
		return nil
	}(events)
	if e == nil {
		return 0, errors.New("no events found with type roundStart")
	}

	nextRound, err := strconv.ParseInt(strconv.Itoa(e.Date), 10, 64)
	if err != nil {
		log.Println("error parsing int", err)
		return 0, nil
	}

	tm := time.Unix(nextRound, 0)
	now := time.Now()
	diff := tm.Sub(now)
	if err != nil {
		log.Println("error substracting now to nextround", err)
		return 0, nil
	}

	return int(diff.Hours() / 24), nil
}
