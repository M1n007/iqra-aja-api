package iqra

import (
	"iqra-aja-api/helpers/databases/mongodb"
	"net/http"
	"strconv"

	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllItems returns a list of all database items to the response.
func GetAllSurah(w http.ResponseWriter, req *http.Request) {
	var size int = 10
	var page int = 1

	vars := mux.Vars(req)
	row := vars["size"]
	param := vars["page"]

	size, _ = strconv.Atoi(row)
	page, _ = strconv.Atoi(param)

	collection := "surah"
	rs, err := mongodb.GetAll(collection, size, page)
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func GetAyatBySurahId(w http.ResponseWriter, req *http.Request) {
	var size int = 10
	var page int = 1

	var collection string = "ayat"

	vars := mux.Vars(req)
	id := vars["id"]
	row := vars["size"]
	param := vars["page"]

	size, _ = strconv.Atoi(row)
	page, _ = strconv.Atoi(param)

	rs, err := mongodb.GetMulti(id, collection, size, page)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}
