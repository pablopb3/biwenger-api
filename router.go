package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	v := "v1"
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/"+v+"/", greet)

	router.HandleFunc("/"+v+"/login", login)

	router.HandleFunc("/"+v+"/getMyPlayers", GetMyPlayers)
	router.HandleFunc("/"+v+"/getPlayerById", GetPlayerById)
	router.HandleFunc("/"+v+"/setLineUp", SetLineUp)

	router.HandleFunc("/"+v+"/sendPlayersToMarket", SendPlayersToMarket)
	router.HandleFunc("/"+v+"/getPlayersInMarket", GetPlayersInMarket)

	router.HandleFunc("/"+v+"/getReceivedOffers", GetReceivedOffers)
	router.HandleFunc("/"+v+"/acceptReceivedOffer", AcceptReceivedOffer)
	router.HandleFunc("/"+v+"/placeOffer", PlaceOffer)

	router.HandleFunc("/"+v+"/getMyMoney", GetMyMoney)
	router.HandleFunc("/"+v+"/getMaxBid", GetMaxBid)
	router.HandleFunc("/"+v+"/getMarketEvolution", GetMarketEvolution)

	router.HandleFunc("/"+v+"/updatePlayersAlias", UpdatePlayersAliasInDb)

	router.HandleFunc("/"+v+"/tweet", Tweet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, SendApiResponseWithMessage("", "Hi from your biwenger api!!"))
}
