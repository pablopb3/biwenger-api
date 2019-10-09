package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"

	"github.com/gin-gonic/gin"
)

const homeURL = "/home"

func getDaysToNextRound(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := biwenger.Request{
			Method:   "GET",
			Endpoint: homeURL,
			Token:    c.GetHeader("authorization"),
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for home", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var home biwenger.Home
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
}

func calculateDays(events []biwenger.Event) (int, error) {
	e := func(events []biwenger.Event) *biwenger.Event {
		for _, event := range events {
			if event.Type == "roundStart" {
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
