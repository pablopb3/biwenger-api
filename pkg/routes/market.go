package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
)

const marketURL string = "/market"
const marketStatsURL string = "/competitions/la-liga/market?interval=day&includeValues=true"

type reqBody struct {
	marketValuePercentatge string `json:"marketValuePercentatge"`
}

func sendPlayers(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody reqBody
		if err := c.BindJSON(&reqBody); err != nil {
			log.Println("error binding reqBody", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		sendToMarket := biwenger.SendToMarket{Type: "team", Price: reqBody.marketValuePercentatge}
		marketJSON, err := json.Marshal(sendToMarket)

		req := biwenger.Request{
			Method:   "POST",
			Endpoint: marketURL,
			Body:     marketJSON,
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for sending players to market", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var statusResponse biwenger.StatusResponse
		err = json.Unmarshal(body, &statusResponse)
		if err != nil {
			log.Println("error unmarshalling response body to statusResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, statusResponse)
	}
}

func marketPlayers(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var playersInMarket []biwenger.PlayerInMarket
		marketResponse, err := getMarket(cli, c)
		if err != nil {
			log.Println("error getting market in marketPlayers", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		for _, sale := range marketResponse.Data.Sales {
			if !isMyPlayer(sale.User.ID) {
				playersInMarket = append(playersInMarket,
					biwenger.PlayerInMarket{
						IdPlayer: sale.Player.ID,
						Price:    sale.Price,
						IdUser:   sale.User.ID,
					},
				)
			}
		}

		c.JSON(http.StatusOK, playersInMarket)
	}
}
func receivedOffers(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var receivedOffers []biwenger.ReceivedOffer
		marketResponse, err := getMarket(cli, c)
		if err != nil {
			log.Println("error getting market in receivedOffers", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		for _, offer := range marketResponse.Data.Offers {
			receivedOffers = append(receivedOffers,
				biwenger.ReceivedOffer{
					IdOffer:  offer.ID,
					IdPlayer: offer.RequestedPlayers[0],
					Amount:   offer.Amount,
					IdUser:   offer.From.ID,
				},
			)
		}

		c.JSON(http.StatusOK, receivedOffers)
	}
}
func myMoney(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		marketResponse, err := getMarket(cli, c)
		if err != nil {
			log.Println("error getting market in myMoney", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, marketResponse.Data.Status.Balance)
	}
}
func myMaxBid(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		marketResponse, err := getMarket(cli, c)
		if err != nil {
			log.Println("error getting market in myMaxBid", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, marketResponse.Data.Status.MaximumBid)
	}
}
func marketEvolution(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := biwenger.Request{
			Method:   "GET",
			Endpoint: marketStatsURL,
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for getting market evolution", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var marketStatsResponse biwenger.MarketStatsResponse
		err = json.Unmarshal(body, &marketStatsResponse)
		if err != nil {
			log.Println("error unmarshalling response body to marketStatsResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, marketStatsResponse.Data.Values)
	}
}
func getMarket(cli biwenger.Client, c *gin.Context) (*biwenger.MarketResponse, error) {
	req := biwenger.Request{
		Method:   "GET",
		Endpoint: marketURL,
	}

	body, err := cli.DoRequest(req)
	if err != nil {
		log.Println("error doing request for getting market", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	var marketResponse biwenger.MarketResponse
	err = json.Unmarshal(body, &marketResponse)
	if err != nil {
		log.Println("error unmarshalling response body to marketResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return &marketResponse, nil
}

func isMyPlayer(userID int) bool {
	p := properties.MustLoadFile("application.properties", properties.UTF8)

	return p.GetInt("userId", 0) == userID
}
