package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/magiconair/properties"
	"net/http"
)


func Tweet(w http.ResponseWriter, r *http.Request) { //TODO unhardcore values

	p := properties.MustLoadFile("application.properties", properties.UTF8)

config := oauth1.NewConfig(p.GetString("twitter.consumerKey", ""), p.GetString("consumerSecret", ""))
	token := oauth1.NewToken(p.GetString("token", ""), p.GetString("tokenSecret", ""))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	client.Statuses.Update("tweet", nil)
	fmt.Fprintf(w, SendApiResponse("tweet ok!"))
}