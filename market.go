package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"net/http"
	"strconv"
)

const marketUrl string = "https://biwenger.as.com/api/v2/market"
const marketStatsUrl string = "https://cf.biwenger.com/api/v2/competitions/la-liga/market?interval=day&includeValues=true"

type SendToMarket struct {
	Type  string `json:"type"`
	Price string    `json:"price"`
}

type BiwengerStatusResponse struct {
	Status int `json:"status"`
	Data   string    `json:"data"`
}


func SendPlayersToMarket(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	price := r.FormValue("price")
	fmt.Println(string(price))
	sendToMarket := SendToMarket{"team", "125"}
	jsonSendToMarket := structToJson(sendToMarket)
	var biwengerResponse = new(BiwengerStatusResponse)
	doRequestAndGetStruct("POST", marketUrl, getDefaultHeaders(r), string(jsonSendToMarket), &biwengerResponse)
	fmt.Fprintf(w, SendApiResponse(biwengerResponse))

}

func GetPlayersInMarket(w http.ResponseWriter, r *http.Request) {
	var playersInMarket []PlayerInMarket
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	for _, sale := range biwengerMarketResponse.Data.Sales {
		if (!IsMyPlayer(sale.User.ID)) {
			playersInMarket = append(playersInMarket, PlayerInMarket{sale.Player.ID, sale.Price, sale.User.ID})
		}
	}
	fmt.Fprintf(w, SendApiResponse(playersInMarket))
}

func GetReceivedOffers(w http.ResponseWriter, r *http.Request) {
	var playersInMarket []ReceivedOffer
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	for _, offer := range biwengerMarketResponse.Data.Offers {
		playersInMarket = append(playersInMarket, ReceivedOffer{offer.ID, offer.RequestedPlayers[0], offer.Amount, offer.From.ID})
	}
	fmt.Fprintf(w, SendApiResponse(&playersInMarket))
}

func GetMyMoney(w http.ResponseWriter, r *http.Request) {
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	fmt.Fprintf(w, SendApiResponse(biwengerMarketResponse.Data.Status.Balance))
}

func GetMaxBid(w http.ResponseWriter, r *http.Request) {
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	fmt.Fprintf(w, SendApiResponse(strconv.Itoa(biwengerMarketResponse.Data.Status.MaximumBid)))
}

func GetMarketEvolution(w http.ResponseWriter, r *http.Request) {
	biwengerMarketEvolutionResponse := getBiwengerMarketEvolutionResponse(r)
	fmt.Fprintf(w, SendApiResponse(biwengerMarketEvolutionResponse.Data.Values))
}

func getBiwengerMarketResponse(r *http.Request) *BiwengerMarketResponse {
	var biwengerMarketResponse = new(BiwengerMarketResponse)
	doRequestAndGetStruct("GET", marketUrl, getDefaultHeaders(r), "", &biwengerMarketResponse)
	return biwengerMarketResponse
}

func getBiwengerMarketEvolutionResponse(r *http.Request) *BiwengerMarketStatsResponse {
	var biwengerMarketStatsResponse = new(BiwengerMarketStatsResponse)
	doRequestAndGetStruct("GET", marketStatsUrl, getDefaultHeaders(r), "", &biwengerMarketStatsResponse)
	return biwengerMarketStatsResponse
}


func IsMyPlayer(userId int) bool {
	p := properties.MustLoadFile("application.properties", properties.UTF8)
	return p.GetInt("userId", 0) == userId

}

type PlayerInMarket struct {
	IdPlayer int `json:"idPlayer"`
	Price 	 int `json:"price"`
	IdUser	 int `json:"idUser"`
}

type ReceivedOffer struct {
	IdOffer	 int `json:"idOffer"`
	IdPlayer int `json:"idPlayer"`
	Ammount	 int `json:"ammount"`
	IdUser	 int `json:"idUser"`
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
			Price int         `json:"price"`
			Until int         `json:"until"`
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