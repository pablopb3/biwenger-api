package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/pablopb3/biwenger-api/pkg/biwenger"
	"github.com/pablopb3/biwenger-api/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	playerAliasPlaceholder = "{playerAlias}"
	getAllPlayersURL       = "https://cf.biwenger.com/api/v2/competitions/la-liga/data?lang=en&score=2" //2 sofascore // TODO param this
	getPlayerURL           = "https://cf.biwenger.com/api/v2/players/la-liga/" + playerAliasPlaceholder + "?fields=*%2Cteam%2Cfitness%2Creports(points%2Chome%2Cevents%2Cstatus(status%2CstatusText)%2Cmatch(*%2Cround%2Chome%2Caway)%2Cstar)%2Cprices%2Ccompetition%2Cseasons%2Cnews%2Cthreads&score=2&lang=en"
)

func player(db repository.Store, cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		alias, err := db.GetAlias(id)
		if err != nil {
			log.Printf("error getting alias for player id %s:\n%s", id, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req := biwenger.Request{
			Method:   "GET",
			Endpoint: strings.Replace(getPlayerURL, playerAliasPlaceholder, alias, 1),
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Printf("error doing request for getting player with alias %s:\n%s", alias, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var player biwenger.Player
		err = json.Unmarshal(body, &player)
		if err != nil {
			log.Println("error unmarshalling response body to player", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, player)
	}
}

func updatePlayers(db repository.Store, cli biwenger.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := biwenger.Request{
			Method:   "GET",
			Endpoint: getAllPlayersURL,
		}

		body, err := cli.DoRequest(req)
		if err != nil {
			log.Println("error doing request for getting all players", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		players := gjson.Get(string(body), "data.players")
		players.ForEach(func(key, value gjson.Result) bool {
			id := gjson.Get(value.String(), "id")
			alias := gjson.Get(value.String(), "slug")
			playerDB := repository.PlayerDB{
				IDPlayer: id.String(),
				Alias:    alias.String(),
			}

			db.Save(playerDB)
			return true // keep iterating
		})

		c.JSON(http.StatusOK, "db successfully updated with players alias!")
	}
}
