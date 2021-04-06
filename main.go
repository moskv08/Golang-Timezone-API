package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moskv08/go-timezone-rocket/models"
)

func main() {

	// Init router
	router := mux.NewRouter()
	router.HandleFunc("/timezone/{zone}", GetTimeByZone).Methods(http.MethodGet)
	router.HandleFunc("/mytimezone", GetMyTimeZone).Methods(http.MethodGet)

	// Start webserver
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetTimeByZone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if zone, ok := params["zone"]; ok {

		tz := models.Timezone{Zone: zone}

		if tz.Time != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(tz.FindProperZone(tz.Zone))
		}
	}
}

func GetMyTimeZone(w http.ResponseWriter, r *http.Request) {
	tz := models.Timezone{}

	if tz.Time != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(tz.FindMyZone())
	}
}
