package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Evrard-Nil/middleware/internal/donneestruct"
)

var layoutHeure = "2006-01-02T15:04:05Z"
var layoutDate = "2006-01-02"

func main() {
	http.HandleFunc("/api/v1/mesures/", mesuresHandler)
	http.HandleFunc("/api/v1/moyennes/", moyennesHandler)
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}

func mesuresHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 6 {
			aeroport := urlParts[4]
			nature := urlParts[5]
			queryValues := r.URL.Query()
			beginDate := queryValues.Get("beginDate")
			beginTime, err1 := time.Parse(layoutHeure, beginDate)
			endDate := queryValues.Get("endDate")
			endTime, err2 := time.Parse(layoutHeure, endDate)
			if err1 != nil || err2 != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("begin : ", beginTime)
			fmt.Printf("end : ", endTime)
			test := donneestruct.MonTest{Nature: nature, Aeroport: aeroport}
			w.WriteHeader(http.StatusFound)
			writeJSON(w, test)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func moyennesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 6 {
			aeroport := urlParts[4]
			nature := urlParts[5]
			queryValues := r.URL.Query()
			date := queryValues.Get("date")
			dateTime, err := time.Parse(layoutDate, date)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("date : ", dateTime)
			test := donneestruct.MonTest{Nature: nature, Aeroport: aeroport}
			w.WriteHeader(http.StatusFound)
			writeJSON(w, test)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
