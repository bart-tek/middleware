package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/gomodule/redigo/redis"
)

var layoutHeure = "2006-01-02T15:04:05Z"
var layoutDate = "2006-01-02"
var redisCli redis.Conn

func main() {
	redisCli = newRedisClient("redis-10932.c1.us-west-2-2.ec2.cloud.redislabs.com:10932", "uutPD4Eh1qkYtGWxiuYvfXE7Ri5N7oPQ")
	http.HandleFunc("/api/v1/mesures/", mesuresHandler)
	http.HandleFunc("/api/v1/moyennes/", moyennesHandler)
	defer redisCli.Close()
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}

// Récupérer un type de données d'un aéroprt entre deux dates
func mesuresHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 6 {
			aeroport := urlParts[4]
			nature := urlParts[5]
			queryValues := r.URL.Query()
			beginDate := queryValues.Get("beginDate")
			beginTime, err1 := time.Parse(layoutHeure, beginDate)
			endDate := queryValues.Get("endDate")
			endTime, err2 := time.Parse(layoutHeure, endDate)
			if err1 == nil && err2 == nil {
				getDataBetweenDates(beginTime, endTime)
				test := donneestruct.MonTest{Nature: nature, Aeroport: aeroport}
				w.WriteHeader(http.StatusFound)
				writeJSON(w, test)
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// Récupérer les moyennes pour chaque type de nature pour une aéroport à une date
func moyennesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 5 {
			aeroport := urlParts[4]
			queryValues := r.URL.Query()
			date := queryValues.Get("date")
			dateTime, err := time.Parse(layoutDate, date)
			if err == nil {
				getDataForOneDate(dateTime)
				test := donneestruct.MonTest{Nature: "null", Aeroport: aeroport}
				w.WriteHeader(http.StatusFound)
				writeJSON(w, test)
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

func getDataBetweenDates(beginTime time.Time, endTime time.Time) {
	fmt.Printf("t1 : ", beginTime.Unix())
	fmt.Printf("t2 : ", endTime.Unix())
	// ZRANGEBYSCORE
}

func getDataForOneDate(dateTime time.Time) {
	fmt.Printf("t : ", dateTime.Unix())
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

func newRedisClient(addr string, pass string) redis.Conn {
	client, err := redis.Dial("tcp", addr, redis.DialPassword(pass))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Succesfully connected to Redis at %s\n", addr)
	}
	return client
}
