package iqra

import (
	"iqra-aja-api/helpers/databases/mongodb"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllItems returns a list of all database items to the response.
func GetAllItems(w http.ResponseWriter, req *http.Request) {
	collection := "surah"
	rs, err := mongodb.GetAll(collection)
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

func GetItem(w http.ResponseWriter, req *http.Request) {
	collection := "ayat"
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := mongodb.GetMulti(id, collection)
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
