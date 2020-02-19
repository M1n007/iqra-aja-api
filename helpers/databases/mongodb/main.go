package mongodb

import (
	Helpers "iqra-aja-api/helpers/utils"
	Models "iqra-aja-api/models"
	"log"
	"strconv"

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

func countData(collection string, find map[string]interface{}) int {
	totalData, _ := collections(collection).Find(find).Count()

	return totalData
}

func countDataAyatBySurah(collection string, newId int) int {
	totalData, _ := collections(collection).Find(bson.M{"sura_id": newId}).Count()

	return totalData
}

// GetAll returns all items from the database.
func GetAll(collection string, row int, param int, find map[string]interface{}) (Models.PaginationResponseSurah, error) {
	var res Models.PaginationResponseSurah
	var parameterPage int = row * (param - 1)
	totalData := countData(collection, find)

	meta := Helpers.GenerateMeta(param, row, totalData)

	if err := collections(collection).Find(find).Limit(row).Skip(parameterPage).All(&res.Data); err != nil {
		res.Err = true
		res.Status = 500
		res.Message = "internal server error"
		return res, nil
	}

	res.Err = false
	res.Status = 200
	res.Message = "success get all data surah"
	res.Meta = meta

	return res, nil
}

// GetOne returns a single item from the database.
func GetMulti(id string, collection string, row int, param int) (Models.PaginationResponseAyat, error) {
	var res Models.PaginationResponseAyat
	newId, _ := strconv.Atoi(id)
	var parameterPage int = row * (param - 1)
	totalData := countDataAyatBySurah(collection, newId)

	meta := Helpers.GenerateMeta(param, row, totalData)

	if err := collections(collection).Find(bson.M{"sura_id": newId}).Limit(row).Skip(parameterPage).All(&res.Data); err != nil {
		res.Err = true
		res.Status = 500
		res.Message = "internal server error"
		return res, nil
	}

	res.Err = false
	res.Status = 200
	res.Message = "success get all data ayat "
	res.Meta = meta

	return res, nil
}
