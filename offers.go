package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const offersUrl string = "https://biwenger.as.com/api/v2/offers/"


func AcceptReceivedOffer(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	offerId := r.FormValue("id")
	url := offersUrl + offerId
	actionToffer := ActionToOffer{"accept"}
	jsonActionToOffer := structToJson(actionToffer)
	var acceptOfferBiwengerResponse = new(AcceptOfferBiwengerResponse)
	doRequestAndGetStruct("PUT", url, getDefaultHeaders(r), string(jsonActionToOffer), &acceptOfferBiwengerResponse)
	fmt.Fprintf(w, string(structToJson(*acceptOfferBiwengerResponse)))

}

func PlaceOffer(w http.ResponseWriter, r *http.Request) {

	placeOfferBody := new(PlaceOfferBody)
	getJsonBody(r, &placeOfferBody)
	jsonPlaceOffer, _ := json.Marshal(placeOfferBody)
	headers := getDefaultHeaders(r)
	placeOfferBiwengerResponse := new(PlaceOfferBiwengerResponse)
	doRequestAndGetStruct("POST", offersUrl, headers, string(jsonPlaceOffer), &placeOfferBiwengerResponse)
	jsonSetLineUpBiwengerResponse, _ := json.Marshal(placeOfferBiwengerResponse)
	fmt.Fprintf(w, string(jsonSetLineUpBiwengerResponse))

}

type ActionToOffer struct {
	Status string `json:"status"`
}

type AcceptOfferBiwengerResponse struct {
	Data struct {
		Amount   int    `json:"amount"`
		Created  int    `json:"created"`
		FromID   int    `json:"fromID"`
		ID       int    `json:"id"`
		Modified int    `json:"modified"`
		Status   string `json:"status"`
		ToID     int    `json:"toID"`
		Type     string `json:"type"`
	} `json:"data"`
	Status int `json:"status"`
}

type PlaceOfferBody struct {
	Amount           int         `json:"amount"`
	RequestedPlayers []int       `json:"requestedPlayers"`
	To               interface{} `json:"to"`
	Type             string      `json:"type"`
}

type PlaceOfferBiwengerResponse struct {
	Data struct {
		Amount  int    `json:"amount"`
		Created int    `json:"created"`
		FromID  int    `json:"fromID"`
		ID      int    `json:"id"`
		Status  string `json:"status"`
		ToID    int    `json:"toID"`
		Type    string `json:"type"`
		Until   int    `json:"until"`
	} `json:"data"`
	Status int `json:"status"`
}