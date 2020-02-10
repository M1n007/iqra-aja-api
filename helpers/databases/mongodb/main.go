package mongodb

import (
	"fmt"
	"log"
	"strconv"

	Models "iqra-aja-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost/iqra_aja")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("iqra_aja")
}

func collections(collection string) *mgo.Collection {
	return db.C(collection)
}

// GetAll returns all items from the database.
func GetAll(collection string) ([]Models.Surah, error) {
	res := []Models.Surah{}

	if err := collections(collection).Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne returns a single item from the database.
func GetMulti(id string, collection string) ([]Models.Ayat, error) {
	res := []Models.Ayat{}
	newId, _ := strconv.Atoi(id)
	fmt.Println(id, collection)

	if err := collections(collection).Find(bson.M{"sura_id": newId}).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Save inserts an item to the database.
func Save(item Models.Surah, collection string) error {
	return collections(collection).Insert(item)
}

// Remove deletes an item from the database
func Remove(id string, collection string) error {
	return collections(collection).Remove(bson.M{"sura_id": id})
}
