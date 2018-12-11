package main

import(
    "net/http"
    "fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"github.com/pablopb3/biwenger-api/dao"
    "strings"
    "encoding/json"
)

const (
	playerAliasMacro = "{playerAlias}"
	getAllPlayersUrl = "https://cf.biwenger.com/api/v2/competitions/la-liga/data?lang=en&score=2" //2 sofascore
	getPlayerUrl = "https://cf.biwenger.com/api/v2/players/la-liga/" + playerAliasMacro + "?fields=*%2Cteam%2Cfitness%2Creports(points%2Chome%2Cevents%2Cstatus(status%2CstatusText)%2Cmatch(*%2Cround%2Chome%2Caway)%2Cstar)%2Cprices%2Ccompetition%2Cseasons%2Cnews%2Cthreads&score=1&lang=en"
) 


func GetPlayerById(w http.ResponseWriter, r *http.Request) {

	//json := GetAllPlayers()

	id := getPlayerIdFromQueryUrl(r)
	playerAlias := dao.GetAliasByPlayerId(id)
	player := new(Player)
	doRequestAndGetStruct("GET", strings.Replace(getPlayerUrl, playerAliasMacro, playerAlias, 1), getPlayersHeaders(), "", &player)
    playerJson, _ := json.Marshal(player)
    fmt.Fprintf(w, string(playerJson))
}

func getPlayerIdFromQueryUrl(r *http.Request) int {

    ids, ok := r.URL.Query()["id"]
    if !ok || len(ids[0]) < 1 {
        panic("Url Param 'key' is missing")
        return 0
    }
    id, _ := strconv.Atoi(ids[0])
    return id
}


func GetAllPlayers() string {
	objectJson := doRequestAndGetJson("GET", getAllPlayersUrl, getPlayersHeaders(), "")
	players := gjson.Get(objectJson, "data.players.15462")
	return players.String()
	
}




func getPlayersHeaders()  map[string]string {

    var m = make(map[string]string)
    m["Referer"] = "https://biwenger.as.com/players"
    m["User-Agent"] = "Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"
    return m
}

type Object struct {
	data interface{}
}



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