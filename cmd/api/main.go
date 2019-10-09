package main

import (
	"flag"
	"log"

	"github.com/charly3pins/biwenger-api/pkg/biwenger"
	"github.com/charly3pins/biwenger-api/pkg/repository"
	"github.com/charly3pins/biwenger-api/pkg/routes"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	uri := flag.String("mongo", "mongodb://mongodb:27017", "mongodb connection address")
	mongodb := flag.String("db", "biwenger", "name of the mongodb database")
	collection := flag.String("collection", "players", "collection for players")
	flag.Parse()

	sess, err := mgo.Dial(*uri)
	if err != nil {
		log.Fatalf("Can't connect to mongo: %s \n", err.Error())
	}
	defer sess.Close()

	col := sess.DB(*mongodb).C(*collection)
	store := repository.Store{C: col}
	cli := biwenger.Client{}

	router := gin.Default()
	routes.NewRouter(router, cli, store)
	log.Fatal(router.Run(":8080"))
}
