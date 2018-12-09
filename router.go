package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/login", login)
    router.HandleFunc("/getMyPlayers", GetMyPlayers)
    router.HandleFunc("/getPlayerById", GetPlayerById)
    router.HandleFunc("/setLineUp", SetLineUp)
    log.Fatal(http.ListenAndServe(":8080", router))
}



