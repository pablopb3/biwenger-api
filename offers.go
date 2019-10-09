package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const offersURL string = "/offers"

func (cli Client) AcceptOffer(c *gin.Context) {
	id := c.Param("id")
	actionToffer := ActionToOffer{"accepted"}

	offerJSON, err := json.Marshal(actionToffer)
	if err != nil {
		log.Println("error marshalling actionToOffer", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := request{
		method:   "PUT",
		endpoint: offersURL + "/" + id,
		body:     offerJSON,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Printf("error doing request for getting accepting offer with id %s:\n%s", id, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var acceptOfferResponse AcceptOfferBiwengerResponse
	err = json.Unmarshal(body, &acceptOfferResponse)
	if err != nil {
		log.Println("error unmarshalling response body to AcceptOfferBiwengerResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, acceptOfferResponse)

}

func (cli Client) PlaceOffer(c *gin.Context) {
	var placeOfferBody PlaceOfferBody
	if err := c.BindJSON(&placeOfferBody); err != nil {
		log.Println("error binding placeOfferBody", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	placeOfferBodyJSON, err := json.Marshal(placeOfferBody)
	if err != nil {
		log.Println("error marshalling placeOfferBody", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := request{
		method:   "POST",
		body:     placeOfferBodyJSON,
		endpoint: offersURL,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for login", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var placeOfferBiwengerResponse PlaceOfferBiwengerResponse
	err = json.Unmarshal(body, &placeOfferBiwengerResponse)
	if err != nil {
		log.Println("error unmarshalling response body to placeOfferBiwengerResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, placeOfferBiwengerResponse)
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
