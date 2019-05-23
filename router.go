package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", greet)

	router.HandleFunc("/login", login)

	router.HandleFunc("/getMyPlayers", GetMyPlayers)
	router.HandleFunc("/updatePlayersAlias", UpdatePlayersAliasInDb)
	router.HandleFunc("/getPlayerById", GetPlayerById)
	router.HandleFunc("/setLineUp", SetLineUp)

	router.HandleFunc("/sendPlayersToMarket", SendPlayersToMarket)
	router.HandleFunc("/getPlayersInMarket", GetPlayersInMarket)

	router.HandleFunc("/getReceivedOffers", GetReceivedOffers)
	router.HandleFunc("/acceptReceivedOffer", AcceptReceivedOffer)
	router.HandleFunc("/placeOffer", PlaceOffer)

	router.HandleFunc("/getMyMoney", GetMyMoney)
	router.HandleFunc("/getMaxBid", GetMaxBid)


	log.Fatal(http.ListenAndServe(":8080", router))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from your biwenger api!!")
}
