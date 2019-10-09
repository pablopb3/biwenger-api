package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"

	"github.com/gin-gonic/gin"
)

const myPlayersURL = "/user?fields=*,lineup(type,playersID),players(*,fitness,team,owner),market(*,-userID),offers,-trophies"
const setLineUpURL = "/user?fields=*"

func myPlayers(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := biwenger.Request{
			Method:   "GET",
			Endpoint: myPlayersURL,
			Token:    c.GetHeader("authorization"),
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for myPlayers", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var currentLineUpResponse biwenger.CurrentLineUpResponse
		err = json.Unmarshal(body, &currentLineUpResponse)
		if err != nil {
			log.Println("error unmarshalling response body to currentLineUpResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var playerIDs []int
		for _, player := range currentLineUpResponse.Data.Players {
			playerIDs = append(playerIDs, player.ID)
		}

		c.JSON(http.StatusOK, playerIDs)
	}
}

func setLineUp(cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lineUp biwenger.LineUp
		if err := c.BindJSON(&lineUp); err != nil {
			log.Println("error binding lineUp", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		lineupJSON, err := json.Marshal(biwenger.StartingEleven{L: lineUp})
		if err != nil {
			log.Println("error marshalling lineUp", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req := biwenger.Request{
			Method:   "PUT",
			Body:     lineupJSON,
			Endpoint: setLineUpURL,
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for setLineUp", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var lineUpResponse biwenger.LineUpResponse
		err = json.Unmarshal(body, &lineUpResponse)
		if err != nil {
			log.Println("error unmarshalling response body to lineUpResponse", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, lineUpResponse)
	}
}
