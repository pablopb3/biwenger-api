package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const baseURL = "https://biwenger.as.com/api/v2"

func main() {
	cli := newClient(baseURL)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/ping", pong)
	v1.POST("/login", cli.login)
	v1.GET("/getDaysToNextRound", cli.getDaysToNextRound)

	// router.HandleFunc("/"+v+"/getDaysToNextRound", GetDaysToNextRound)
	// router.HandleFunc("/"+v+"/getMyPlayers", GetMyPlayers)
	// router.HandleFunc("/"+v+"/getPlayerById", GetPlayerById)
	// router.HandleFunc("/"+v+"/setLineUp", SetLineUp)

	// router.HandleFunc("/"+v+"/sendPlayersToMarket", SendPlayersToMarket)
	// router.HandleFunc("/"+v+"/getPlayersInMarket", GetPlayersInMarket)

	// router.HandleFunc("/"+v+"/getReceivedOffers", GetReceivedOffers)
	// router.HandleFunc("/"+v+"/acceptReceivedOffer", AcceptReceivedOffer)
	// router.HandleFunc("/"+v+"/placeOffer", PlaceOffer)

	// router.HandleFunc("/"+v+"/getMyMoney", GetMyMoney)
	// router.HandleFunc("/"+v+"/getMaxBid", GetMaxBid)
	// router.HandleFunc("/"+v+"/getMarketEvolution", GetMarketEvolution)

	// router.HandleFunc("/"+v+"/updatePlayersAlias", UpdatePlayersAliasInDb)

	// router.HandleFunc("/"+v+"/tweet", Tweet)

	log.Fatal(router.Run(":8080"))
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
