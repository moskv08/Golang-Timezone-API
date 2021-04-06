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

	// Start webserver
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetTimeByZone(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// var err error
	if zone, ok := params["zone"]; ok {
		// userID, err = strconv.Atoi(val)
		// if err != nil {
		// 		w.WriteHeader(http.StatusInternalServerError)
		// 		w.Write([]byte(`{"message": "need a number"}`))
		// 		return
		// }
		tz := models.FindProperZone(zone)

		if tz.Time != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(tz)
		}

	}

}
