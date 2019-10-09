package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

const baseURL = "https://biwenger.as.com/api/v2"

func main() {
	uri := flag.String("mongo", "mongodb://mongodb:27017", "mongodb connection address")
	mongodb := flag.String("db", "biwenger", "name of the mongodb database")
	collection := flag.String("collection", "players", "collection for players")
	flag.Parse()

	sess, err := mgo.Dial(*uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo: %s \n", err.Error())
	}
	defer sess.Close()

	col := sess.DB(*mongodb).C(*collection)
	s := Store{c: col}
	cli := Client{apiURL: baseURL}

	router := gin.Default()

	router.GET("/greet", greet)

	router.POST("/login", cli.login)

	squad := router.Group("/squad")
	squad.GET("", cli.myPlayers)
	squad.PUT("/lineup", cli.setLineUp)
	squad.GET("/money", cli.GetMyMoney)
	squad.GET("/maxbid", cli.GetMaxBid)

	market := router.Group("/market")
	market.GET("/evolution", cli.GetMarketEvolution)
	market.PUT("/squad", cli.SendPlayersToMarket)
	market.PUT("/squad/:playerID", nil) // For the future
	market.GET("/players", cli.GetPlayersInMarket)
	market.POST("/offers", cli.PlaceOffer)
	market.GET("/offers", cli.GetReceivedOffers)
	market.PUT("/offers/:id", cli.AcceptOffer)

	router.GET("/players/:id", cli.getPlayer(s))
	router.PUT("/players", cli.updatePlayers(s))

	router.GET("/rounds/next", cli.getDaysToNextRound)

	log.Fatal(router.Run(":8080"))
}

func greet(c *gin.Context) {
	c.JSON(http.StatusOK, "Hi from your biwenger api!!")
}
