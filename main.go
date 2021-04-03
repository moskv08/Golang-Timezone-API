package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Init router
	router := mux.NewRouter()
	router.HandleFunc("/timezone", GetTimeByZone).Methods(http.MethodGet)

	// Start webserver
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetTimeByZone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}
