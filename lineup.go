package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const myPlayersURL = "/user?fields=*,lineup(type,playersID),players(*,fitness,team,owner),market(*,-userID),offers,-trophies"
const setLineUpURL = "/user?fields=*"

func (cli Client) myPlayers(c *gin.Context) {
	req := request{
		method:   "GET",
		endpoint: myPlayersURL,
		token:    c.GetHeader("authorization"),
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for myPlayers", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var myPlayers MyPlayers
	err = json.Unmarshal(body, &myPlayers)
	if err != nil {
		log.Println("error unmarshalling response body to myPlayers", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var playerIDs []int
	for _, player := range myPlayers.Data.Players {
		playerIDs = append(playerIDs, player.ID)
	}

	c.JSON(http.StatusOK, playerIDs)
}

func (cli Client) setLineUp(c *gin.Context) {
	var lineUp LineUp
	if err := c.BindJSON(&lineUp); err != nil {
		log.Println("error binding lineUp", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	lineupJSON, err := json.Marshal(StartingEleven{L: lineUp})
	if err != nil {
		log.Println("error marshalling lineUp", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := request{
		method:   "PUT",
		body:     lineupJSON,
		endpoint: setLineUpURL,
	}

	body, err := cli.doRequest(req)
	if err != nil {
		log.Println("error doing request for setLineUp", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var setLineUp SetLineUp
	err = json.Unmarshal(body, &setLineUp)
	if err != nil {
		log.Println("error unmarshalling response body to setLineUp", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, setLineUp)
}

type MyPlayers struct {
	Status int        `json:"status"`
	Data   LineUpData `json:"data"`
}

type LineUpData struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Group      string        `json:"group"`
	Icon       string        `json:"icon"`
	Points     int           `json:"points"`
	Balance    int           `json:"balance"`
	JoinDate   int           `json:"joinDate"`
	LineupDate int           `json:"lineupDate"`
	LineUp     LineUp        `json:"lineup"`
	Market     []interface{} `json:"market"`
	Players    []PlayerBase  `json:"players"`
	Offers     []interface{} `json:"offers"`
}

type SetLineUp struct {
	Status int            `json:"status"`
	Data   LineUpBaseData `json:"data"`
}

type LineUpBaseData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Group      string `json:"group"`
	Icon       string `json:"icon"`
	Points     int    `json:"points"`
	Balance    int    `json:"balance"`
	JoinDate   int    `json:"joinDate"`
	LineupDate int    `json:"lineupDate"`
}

type StartingEleven struct {
	L LineUp `json:"lineup"`
}

type LineUp struct {
	Formation string `json:"type"`
	PlayerIds []int  `json:"playersID"`
}

type Owner struct {
	Date  int `json:"date"`
	Price int `json:"price"`
}

type PlayerBase struct {
	ID    int   `json:"id"`
	Owner Owner `json:"owner"`
}
