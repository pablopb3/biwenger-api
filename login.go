package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

type Login struct {
	Token Token `json:"login"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const loginUrl string = "https://biwenger.as.com/api/v2/auth/login"

func login(w http.ResponseWriter, r *http.Request) {

	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	headers := getLoginHeaders()

	token := new(Token)
	doRequestAndGetStruct("POST", loginUrl, headers, string(userJson), token)
	login := Login{*token}
	fmt.Fprintf(w, SendApiResponse(login))
}

func getLoginHeaders() map[string]string {

	var m = make(map[string]string)
	m["Content-Type"] = "application/json"
	return m
}
