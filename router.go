package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", sayHi)
	router.HandleFunc("/login", login)
	router.HandleFunc("/getMyPlayers", GetMyPlayers)
	router.HandleFunc("/updatePlayersAlias", UpdatePlayersAliasInDb)
	router.HandleFunc("/getPlayerById", GetPlayerById)
	router.HandleFunc("/setLineUp", SetLineUp)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi from your biwenger api!!")
}
