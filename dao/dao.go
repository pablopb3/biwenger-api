package dao

import(
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	SERVER = "mongodb://localhost:27017"
	DATABASE = "biwenger"
	COLLECTION = "players"
)

type PlayerAlias struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	IdPlayer int `bson:"id_player" json:"id_player"`
	Alias string `bson:"alias" json:"alias"`
}

func GetAliasByPlayerId(id int) string {

	var db *mgo.Database
	session, err := mgo.Dial(SERVER)
	if err != nil {
		panic(err)
	}

	var player PlayerAlias
	db = session.DB(DATABASE)
	db.C(COLLECTION).Find(bson.M{"id_player": id}).One(&player)
	return player.Alias
}