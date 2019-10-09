package routes

import (
	"net/http"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"
	"github.com/pablopb3/biwenger-api/pkg/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, cli biwenger.Client, store repository.Store) {
	router.GET("/greet", greet)

	router.POST("/login", login(cli))

	squad := router.Group("/squad")
	squad.GET("", myPlayers(cli))
	squad.PUT("/lineup", setLineUp(cli))
	squad.GET("/money", myMoney(cli))
	squad.GET("/maxbid", myMaxBid(cli))

	market := router.Group("/market")
	market.GET("/evolution", marketEvolution(cli))
	market.PUT("/squad", sendPlayers(cli))
	//market.PUT("/squad/:playerID", nil) // For the future
	market.GET("/players", marketPlayers(cli))
	market.POST("/offers", placeOffer(cli))
	market.GET("/offers", receivedOffers(cli))
	market.PUT("/offers/:id", acceptOffer(cli))

	router.GET("/players/:id", player(store, cli))
	router.PUT("/players", updatePlayers(store, cli))

	router.GET("/rounds/next", getDaysToNextRound(cli))
}

func greet(c *gin.Context) {
	c.JSON(http.StatusOK, "Hi from your biwenger api!")
}
