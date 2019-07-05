package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
)


func Tweet(w http.ResponseWriter, r *http.Request) { //TODO unhardcore values
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("token", "tokenSecret")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	client.Statuses.Update("tweet", nil)
	fmt.Fprintf(w, SendApiResponse("tweet ok!"))
}