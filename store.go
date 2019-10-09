package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Store struct {
	c *mgo.Collection
}

type PlayerDB struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	IDPlayer string        `bson:"id_player" json:"id_player"`
	Alias    string        `bson:"alias" json:"alias"`
}

func (s Store) getAlias(id string) (string, error) {
	var player PlayerDB
	err := s.c.Find(bson.M{"_id": id}).One(&player)
	if err != nil {
		log.Printf("error finding player id %s:\n%s", id, err)
		return "", err
	}

	return player.Alias, nil
}

func (s Store) save(player PlayerDB) error {
	return s.c.Insert(player)
}
