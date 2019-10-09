package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
)

const marketURL string = "/market"
const marketStatsURL string = "/competitions/la-liga/market?interval=day&includeValues=true"

type PlayerInMarket struct {
	IdPlayer int `json:"idPlayer"`
	Price    int `json:"price"`
	IdUser   int `json:"idUser"`
}

type ReceivedOffer struct {
	IdOffer  int `json:"idOffer"`
	IdPlayer int `json:"idPlayer"`
	Ammount  int `json:"ammount"`
	IdUser   int `json:"idUser"`
}

type BiwengerMarketResponse struct {
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
			Price int `json:"price"`
			Until int `json:"until"`
			User  struct {
				ID int `json:"id"`
			} `json:"user"`
		} `json:"sales"`
		Status struct {
			Balance    int `json:"balance"`
			MaximumBid int `json:"maximumBid"`
		} `json:"status"`
	} `json:"data"`
	Status int `json:"status"`
}

type BiwengerMarketStatsResponse struct {
	Status int `json:"status"`
	Data   struct {
		Competition struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			Sport    string `json:"sport"`
			Currency string `json:"currency"`
			Country  string `json:"country"`
			Enabled  bool   `json:"enabled"`
			Type     string `json:"type"`
		} `json:"competition"`
		Values [][]int `json:"values"`
		Ups    []struct {
			OldPrice       int    `json:"oldPrice"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			Position       int    `json:"position"`
			Price          int    `json:"price"`
			FantasyPrice   int    `json:"fantasyPrice"`
			Country        string `json:"country"`
			Birthday       int    `json:"birthday"`
			Status         string `json:"status"`
			PriceIncrement int    `json:"priceIncrement"`
			Difference     int    `json:"difference"`
			Team           struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"team"`
			Fitness          []int  `json:"fitness"`
			Points           int    `json:"points"`
			PlayedHome       int    `json:"playedHome"`
			PlayedAway       int    `json:"playedAway"`
			PointsHome       int    `json:"pointsHome"`
			PointsAway       int    `json:"pointsAway"`
			PointsLastSeason int    `json:"pointsLastSeason"`
			StatusText       string `json:"statusText,omitempty"`
		} `json:"ups"`
		Downs []struct {
			OldPrice       int    `json:"oldPrice"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			Slug           string `json:"slug"`
			Position       int    `json:"position"`
			Price          int    `json:"price"`
			FantasyPrice   int    `json:"fantasyPrice"`
			Country        string `json:"country"`
			Birthday       int    `json:"birthday"`
			Status         string `json:"status"`
			PriceIncrement int    `json:"priceIncrement"`
			Difference     int    `json:"difference"`
			Team           struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"team"`
			Fitness          []interface{} `json:"fitness"`
			Points           int           `json:"points"`
			PlayedHome       int           `json:"playedHome"`
			PlayedAway       int           `json:"playedAway"`
			PointsHome       int           `json:"pointsHome"`
			PointsAway       int           `json:"pointsAway"`
			PointsLastSeason int           `json:"pointsLastSeason"`
			StatusText       string        `json:"statusText,omitempty"`
		} `json:"downs"`
	} `json:"data"`
}

type SendToMarket struct {
	Type  string `json:"type"`
	Price string `json:"price"`
}

type BiwengerStatusResponse struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func (cli Client) SendPlayersToMarket(c *gin.Context) {
	price := c.Query("price") // change to marketValuePercentatge from the body
	sendToMarket := SendToMarket{Type: "team", Price: price}
	marketJSON, err := json.Marshal(sendToMarket)

	req := request{
		method:   "POST",
		endpoint: marketURL,
		body:     marketJSON,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for sending players to market", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var biwengerStatusResponse BiwengerStatusResponse
	err = json.Unmarshal(body, &biwengerStatusResponse)
	if err != nil {
		log.Println("error unmarshalling response body to biwengerStatusResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, biwengerStatusResponse)

}

func (cli Client) GetPlayersInMarket(c *gin.Context) {
	var playersInMarket []PlayerInMarket
	biwengerMarketResponse, err := cli.getMarket(c)
	if err != nil {
		log.Println("error getting market in GetPlayersInMarket", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, sale := range biwengerMarketResponse.Data.Sales {
		if !isMyPlayer(sale.User.ID) {
			playersInMarket = append(playersInMarket, PlayerInMarket{sale.Player.ID, sale.Price, sale.User.ID})
		}
	}

	c.JSON(http.StatusOK, playersInMarket)
}

func (cli Client) GetReceivedOffers(c *gin.Context) {
	var receivedOffers []ReceivedOffer
	biwengerMarketResponse, err := cli.getMarket(c)
	if err != nil {
		log.Println("error getting market in GetReceivedOffers", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, offer := range biwengerMarketResponse.Data.Offers {
		receivedOffers = append(receivedOffers, ReceivedOffer{offer.ID, offer.RequestedPlayers[0], offer.Amount, offer.From.ID})
	}

	c.JSON(http.StatusOK, receivedOffers)
}

func (cli Client) GetMyMoney(c *gin.Context) {
	biwengerMarketResponse, err := cli.getMarket(c)
	if err != nil {
		log.Println("error getting market in GetReceivedOffers", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, biwengerMarketResponse.Data.Status.Balance)
}

func (cli Client) GetMaxBid(c *gin.Context) {
	biwengerMarketResponse, err := cli.getMarket(c)
	if err != nil {
		log.Println("error getting market in GetReceivedOffers", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, biwengerMarketResponse.Data.Status.MaximumBid)
}

func (cli Client) GetMarketEvolution(c *gin.Context) {
	req := request{
		method:   "GET",
		endpoint: marketStatsURL,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for getting market", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var biwengerMarketStatsResponse BiwengerMarketStatsResponse
	err = json.Unmarshal(body, &biwengerMarketStatsResponse)
	if err != nil {
		log.Println("error unmarshalling response body to biwengerMarketStatsResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, biwengerMarketStatsResponse.Data.Values)
}

func (cli Client) getMarket(c *gin.Context) (*BiwengerMarketResponse, error) {
	req := request{
		method:   "GET",
		endpoint: marketURL,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for getting market", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	var biwengerMarketResponse BiwengerMarketResponse
	err = json.Unmarshal(body, &biwengerMarketResponse)
	if err != nil {
		log.Println("error unmarshalling response body to biwengerMarketResponse", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return &biwengerMarketResponse, nil
}

func isMyPlayer(userId int) bool {
	p := properties.MustLoadFile("application.properties", properties.UTF8)

	return p.GetInt("userId", 0) == userId
}
