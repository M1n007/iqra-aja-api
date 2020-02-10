package server

import (
	"iqra-aja-api/modules/iqra-aja"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init Server
func Init() {
	r := mux.NewRouter()

	//Routes
	r.HandleFunc("/api/surah", iqra.GetAllItems).Methods("GET")
	r.HandleFunc("/api/surah/{id}", iqra.GetItem).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
