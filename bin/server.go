package server

import (
	"iqra-aja-api/modules/iqra-aja"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init Server
func Init() {
	r := mux.NewRouter().StrictSlash(false)

	//Routes
	r.HandleFunc("/api/surah", iqra.GetAllSurah).Queries("page", "{page}", "size", "{size}", "search", "{search}").Methods("GET")
	r.HandleFunc("/api/surah/{id}", iqra.GetAyatBySurahId).Queries("page", "{page:[0-9]+}", "size", "{size:[0-9]+}").Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
