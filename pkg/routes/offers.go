package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"

	"github.com/gin-gonic/gin"
)

const offersURL string = "/offers"

func acceptOffer(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		actionToffer := biwenger.ActionToOffer{Status: "accepted"}
		offerJSON, err := json.Marshal(actionToffer)
		if err != nil {
			log.Println("error marshalling actionToOffer", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		id := c.Param("id")
		req := biwenger.Request{
			Method:   "PUT",
			Endpoint: offersURL + "/" + id,
			Body:     offerJSON,
			Token:    c.GetHeader("authorization"),
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Printf("error doing request for getting accepting offer with id %s:\n%s", id, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var acceptOfferResponse biwenger.AcceptOfferResponse
		err = json.Unmarshal(body, &acceptOfferResponse)
		if err != nil {
			log.Println("error unmarshalling response body to acceptOfferResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, acceptOfferResponse)
	}
}

func placeOffer(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var placeOfferBody biwenger.PlaceOfferBody
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

		req := biwenger.Request{
			Method:   "POST",
			Body:     placeOfferBodyJSON,
			Endpoint: offersURL,
			Token:    c.GetHeader("authorization"),
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for login", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var placeOfferResponse biwenger.PlaceOfferResponse
		err = json.Unmarshal(body, &placeOfferResponse)
		if err != nil {
			log.Println("error unmarshalling response body to placeOfferResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, placeOfferResponse)
	}
}
