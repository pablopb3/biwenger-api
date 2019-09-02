package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const getHomeUrl = "https://biwenger.as.com/api/v2/home"

func GetDaysToNextRound(w http.ResponseWriter, r *http.Request) {

	headers := getDefaultHeaders(r)
	homeBiwengerResponse := new(Home)
	doRequestAndGetStruct("GET", getHomeUrl, headers, "", &homeBiwengerResponse)
	daysToNextRound := GetDaysToNextRoundFromHomeResponse(homeBiwengerResponse.Data.Events)
	fmt.Fprintf(w, SendApiResponse(daysToNextRound))
}

func GetDaysToNextRoundFromHomeResponse(events []Event) int {
	nextEvent := new(Event)
	for _, event := range events {
		if event.Type == "roundStart" {
			nextEvent = &event
			break
		}
	}
	nextRound, err := strconv.ParseInt(strconv.Itoa(nextEvent.Date), 10, 64)
	tm := time.Unix(nextRound, 0)
	fmt.Println(tm)
	now := time.Now()
	fmt.Println(now)
	diff := tm.Sub(now)
	if err != nil {
		panic(err)
	}
	return int(diff.Hours() / 24)
}

type Home struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}
type From struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Content struct {
	Type   string `json:"type"`
	Player int    `json:"player"`
	From   From   `json:"from,omitempty"`
	Team   Team   `json:"team,omitempty"`
}
type Board struct {
	Type    string      `json:"type"`
	Title   string      `json:"title"`
	Content []Content   `json:"content"`
	Date    int         `json:"date"`
	Fixed   int         `json:"fixed"`
	Author  interface{} `json:"author"`
}
type League struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Icon         string  `json:"icon"`
	Type         string  `json:"type"`
	Mode         string  `json:"mode"`
	Board        []Board `json:"board"`
	Participants int     `json:"participants"`
}
type Status struct {
	Points  int `json:"points"`
	Balance int `json:"balance"`
	Offers  int `json:"offers"`
	Bids    int `json:"bids"`
}
type HomeUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Group    string `json:"group"`
	Position int    `json:"position"`
	Type     string `json:"type"`
	Status   Status `json:"status"`
}
type Round struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Short string `json:"short"`
}
type Event struct {
	Type  string `json:"type"`
	Date  int    `json:"date"`
	Round Round  `json:"round"`
}
type Account struct {
	UnreadMessages bool `json:"unreadMessages"`
}
type Data struct {
	League      League   `json:"league"`
	User        HomeUser `json:"user"`
	Competition string   `json:"competition"`
	Events      []Event  `json:"events"`
	Account     Account  `json:"account"`
}
