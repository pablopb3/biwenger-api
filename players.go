package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	playerAliasPlaceholder = "{playerAlias}"
	getAllPlayersURL       = "https://cf.biwenger.com/api/v2/competitions/la-liga/data?lang=en&score=2" //2 sofascore // TODO param this
	getPlayerURL           = "https://cf.biwenger.com/api/v2/players/la-liga/" + playerAliasPlaceholder + "?fields=*%2Cteam%2Cfitness%2Creports(points%2Chome%2Cevents%2Cstatus(status%2CstatusText)%2Cmatch(*%2Cround%2Chome%2Caway)%2Cstar)%2Cprices%2Ccompetition%2Cseasons%2Cnews%2Cthreads&score=2&lang=en"
)

type Player struct {
	Status int `json:"status"`
	Data   struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		Position       int    `json:"position"`
		Price          int    `json:"price"`
		FantasyPrice   int    `json:"fantasyPrice"`
		Status         string `json:"status"`
		PriceIncrement int    `json:"priceIncrement"`
		Competition    struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			Sport    string `json:"sport"`
			Currency string `json:"currency"`
			Country  string `json:"country"`
			Enabled  bool   `json:"enabled"`
			Type     string `json:"type"`
		} `json:"competition"`
		Team struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Slug      string `json:"slug"`
			NextMatch struct {
				ID     int    `json:"id"`
				Date   int    `json:"date"`
				Status string `json:"status"`
				Round  struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Short  string `json:"short"`
					Season struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
						Slug string `json:"slug"`
					} `json:"season"`
				} `json:"round"`
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
				Home struct {
					ID      int           `json:"id"`
					Name    string        `json:"name"`
					Slug    string        `json:"slug"`
					Score   interface{}   `json:"score"`
					Reports []interface{} `json:"reports"`
				} `json:"home"`
				Away struct {
					ID      int           `json:"id"`
					Name    string        `json:"name"`
					Slug    string        `json:"slug"`
					Score   interface{}   `json:"score"`
					Reports []interface{} `json:"reports"`
				} `json:"away"`
				Preview interface{} `json:"preview"`
				Summary interface{} `json:"summary"`
			} `json:"nextMatch"`
		} `json:"team"`
		Reports []struct {
			Match struct {
				ID     int    `json:"id"`
				Date   int    `json:"date"`
				Status string `json:"status"`
				Round  struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Short string `json:"short"`
				} `json:"round"`
				Home struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Slug  string `json:"slug"`
					Score int    `json:"score"`
				} `json:"home"`
				Away struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Slug  string `json:"slug"`
					Score int    `json:"score"`
				} `json:"away"`
			} `json:"match"`
			Events []struct {
				Type     int `json:"type"`
				Metadata int `json:"metadata"`
			} `json:"events"`
			Points int  `json:"points"`
			Star   bool `json:"star,omitempty"`
		} `json:"reports"`
		Prices  [][]int `json:"prices"`
		Seasons []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Competition struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				Sport    string `json:"sport"`
				Currency string `json:"currency"`
				Country  string `json:"country"`
				Enabled  bool   `json:"enabled"`
				Type     string `json:"type"`
			} `json:"competition,omitempty"`
			Player struct {
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"player,omitempty"`
		} `json:"seasons"`
		News []struct {
			ID       int    `json:"id"`
			Source   string `json:"source"`
			Title    string `json:"title"`
			URL      string `json:"url"`
			Date     int    `json:"date"`
			Photo    string `json:"photo"`
			Video    string `json:"video"`
			Comments int    `json:"comments"`
		} `json:"news"`
		Threads []struct {
			ID       int    `json:"id"`
			Source   string `json:"source"`
			Title    string `json:"title"`
			URL      string `json:"url"`
			Date     int    `json:"date"`
			Photo    string `json:"photo"`
			Video    string `json:"video"`
			Comments int    `json:"comments"`
		} `json:"threads"`
		Fitness          []int `json:"fitness"`
		Points           int   `json:"points"`
		PlayedHome       int   `json:"playedHome"`
		PlayedAway       int   `json:"playedAway"`
		PointsHome       int   `json:"pointsHome"`
		PointsAway       int   `json:"pointsAway"`
		PointsLastSeason int   `json:"pointsLastSeason"`
		ScoreID          int   `json:"scoreID"`
	} `json:"data"`
}

type PlayersAliasInfo struct {
	NumID struct {
		ID               int           `json:"id"`
		Name             string        `json:"name"`
		Slug             string        `json:"slug"`
		TeamID           int           `json:"teamID"`
		Position         int           `json:"position"`
		Price            int           `json:"price"`
		FantasyPrice     int           `json:"fantasyPrice"`
		Status           string        `json:"status"`
		PriceIncrement   int           `json:"priceIncrement"`
		Fitness          []interface{} `json:"fitness"`
		Points           int           `json:"points"`
		PlayedHome       int           `json:"playedHome"`
		PlayedAway       int           `json:"playedAway"`
		PointsHome       int           `json:"pointsHome"`
		PointsAway       int           `json:"pointsAway"`
		PointsLastSeason int           `json:"pointsLastSeason"`
	} `json:"id"`
}

func (cli Client) getPlayer(db Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		alias, err := db.getAlias(id)
		if err != nil {
			log.Printf("error getting alias for player id %s:\n%s", id, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		req := request{
			method:   "GET",
			endpoint: strings.Replace(getPlayerURL, playerAliasPlaceholder, alias, 1),
		}

		body, err := cli.doRequest(req)
		if err != nil {
			log.Printf("error doing request for getting player with alias %s:\n%s", alias, err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var player Player
		err = json.Unmarshal(body, &player)
		if err != nil {
			log.Println("error unmarshalling response body to player", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, player)
	}
}

func (cli Client) updatePlayers(db Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := request{
			method:   "GET",
			endpoint: getAllPlayersURL,
		}

		body, err := cli.doRequest(req)
		if err != nil {
			log.Println("error doing request for getting all players", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		players := gjson.Get(string(body), "data.players")
		players.ForEach(func(key, value gjson.Result) bool {
			id := gjson.Get(value.String(), "id")
			alias := gjson.Get(value.String(), "slug")
			playerDB := PlayerDB{
				IDPlayer: id.String(),
				Alias:    alias.String(),
			}

			db.save(playerDB)
			return true // keep iterating
		})

		c.JSON(http.StatusOK, "db successfully updated with players alias!")
	}
}
