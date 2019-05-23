package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"net/http"
	"strconv"
)

const marketUrl string = "https://biwenger.as.com/api/v2/market"

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
	fmt.Fprintf(w, string(structToJson(*biwengerResponse)))

}

func GetPlayersInMarket(w http.ResponseWriter, r *http.Request) {
	var playersInMarket []PlayerInMarket
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	for _, sale := range biwengerMarketResponse.Data.Sales {
		if (!IsMyPlayer(sale.User.ID)) {
			playersInMarket = append(playersInMarket, PlayerInMarket{sale.Player.ID, sale.Price, sale.User.ID})
		}
	}
	fmt.Fprintf(w, string(structToJson(&playersInMarket)))
}

func GetReceivedOffers(w http.ResponseWriter, r *http.Request) {
	var playersInMarket []ReceivedOffer
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	for _, offer := range biwengerMarketResponse.Data.Offers {
		playersInMarket = append(playersInMarket, ReceivedOffer{offer.ID, offer.RequestedPlayers[0], offer.Amount, offer.From.ID})
	}
	fmt.Fprintf(w, string(structToJson(&playersInMarket)))
}

func GetMyMoney(w http.ResponseWriter, r *http.Request) {
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	fmt.Fprintf(w, strconv.Itoa(biwengerMarketResponse.Data.Status.Balance))
}

func GetMaxBid(w http.ResponseWriter, r *http.Request) {
	biwengerMarketResponse := getBiwengerMarketResponse(r)
	fmt.Fprintf(w, strconv.Itoa(biwengerMarketResponse.Data.Status.MaximumBid))
}

func getBiwengerMarketResponse(r *http.Request) *BiwengerMarketResponse {
	var biwengerMarketResponse = new(BiwengerMarketResponse)
	doRequestAndGetStruct("GET", marketUrl, getDefaultHeaders(r), "", &biwengerMarketResponse)
	return biwengerMarketResponse
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