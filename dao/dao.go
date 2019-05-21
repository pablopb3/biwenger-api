package dao

import (
	"encoding/json"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	SERVER     = "mongodb://mongodb:27017"
	DATABASE   = "biwenger"
	COLLECTION = "players"
)

type PlayerAlias struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	IdPlayer int           `bson:"id_player" json:"id_player"`
	Alias    string        `bson:"alias" json:"alias"`
}

func GetAliasByPlayerId(id int) string {

	var db *mgo.Database
	session, err := mgo.Dial(SERVER)
	if err != nil {
		panic(err)
	}

	var player PlayerAlias
	db = session.DB(DATABASE)
	db.C(COLLECTION).Find(bson.M{"_id": id}).One(&player)
	return player.Alias
}

func SavePlayerAlias(playerIdAliasMap PlayerIdAliasMap) {

	var db *mgo.Database
	session, err := mgo.Dial(SERVER)
	if err != nil {
		panic(err)
	}
	db = session.DB(DATABASE)
	db.C(COLLECTION).Insert(playerIdAliasMap)
	playerJson, _ := json.Marshal(playerIdAliasMap)
	fmt.Printf("Player inserted in db: " + string(playerJson))
}

type PlayerIdAliasMap struct {
	ID    int    `bson:"_id" json:"id"`
	Alias string `bson:"alias" json:"alias"`
}
