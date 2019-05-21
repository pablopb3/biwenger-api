package main

import (
	"fmt"
	"net/http"
)

const marketUrl string = "https://biwenger.as.com/api/v2/market"

type SendToMarket struct {
	Type  string `json:"type"`
	Price string    `json:"price"`
}

type BiwengerStatusResponse struct {
	Status int `json:"status"`
	Data   string    `json:"data"`
}

type PlayersInMarket struct {
	Data struct {
		Loans  []interface{} `json:"loans"`
		Offers []struct {
			Amount  int `json:"amount"`
			Created int `json:"created"`
			From    struct {
				Icon string `json:"icon"`
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"from"`
			ID               int    `json:"id"`
			RequestedPlayers []int  `json:"requestedPlayers"`
			Status           string `json:"status"`
			To               struct {
				Icon string `json:"icon"`
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"to"`
			Type  string `json:"type"`
			Until int    `json:"until"`
		} `json:"offers"`
		Sales []struct {
			Date   int `json:"date"`
			Player struct {
				ID int `json:"id"`
			} `json:"player"`
			Price int         `json:"price"`
			Until int         `json:"until"`
			User  interface{} `json:"user"`
		} `json:"sales"`
		Status struct {
			Balance    int `json:"balance"`
			MaximumBid int `json:"maximumBid"`
		} `json:"status"`
	} `json:"data"`
	Status int `json:"status"`
}


func SendPlayersToMarket(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	price := r.FormValue("price")
	fmt.Println(string(price))
	sendToMarket := SendToMarket{"team", "125"}
	jsonSendToMarket := structToJson(sendToMarket)
	var biwengerResponse = new(BiwengerStatusResponse)
	doRequestAndGetStruct("POST", marketUrl, getDefaultHeaders(r), string(jsonSendToMarket), &biwengerResponse)
	fmt.Fprintf(w, string(structToJson(*biwengerResponse)))

}

func GetPlayersInMarket(w http.ResponseWriter, r *http.Request) {

	var playersInMarket = new(PlayersInMarket)
	doRequestAndGetStruct("GET", marketUrl, getDefaultHeaders(r), "", &playersInMarket)
	fmt.Fprintf(w, string(structToJson(*playersInMarket)))

}