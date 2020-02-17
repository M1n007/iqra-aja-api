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

func countData(collection string) int {
	totalData, _ := collections(collection).Count()

	return totalData
}

// GetAll returns all items from the database.
func GetAll(collection string, row int, param int) (Models.PaginationResponseSurah, error) {
	var res Models.PaginationResponseSurah
	var parameterPage int = row * (param - 1)

	totalData := countData(collection)

	meta := bson.M{
		"page":       param,
		"size":       row,
		"total_page": float64(totalData) / float64(row),
		"total_data": totalData,
	}

	if err := collections(collection).Find(nil).Limit(row).Skip(parameterPage).All(&res.Data); err != nil {
		res.Err = true
		res.Status = 500
		res.Message = "internal server error"
		return res, nil
	}

	fmt.Print(meta)

	res.Err = false
	res.Status = 200
	res.Message = "success get all data surah"
	// res.Meta = &meta;

	return res, nil
}

// GetOne returns a single item from the database.
func GetMulti(id string, collection string, row int, param int) ([]Models.Ayat, error) {
	res := []Models.Ayat{}
	newId, _ := strconv.Atoi(id)
	var parameterPage int = row * (param - 1)

	if err := collections(collection).Find(bson.M{"sura_id": newId}).Limit(row).Skip(parameterPage).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}
