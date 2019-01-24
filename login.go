package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Token struct {
	Token string
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

	fmt.Fprintf(w, string(structToJson(&token)))
}

func getLoginHeaders() map[string]string {

	var m = make(map[string]string)
	m["Content-Type"] = "application/json"
	return m
}
