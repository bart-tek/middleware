package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	http.HandleFunc("/api/v1/measures/", measuresHandler)
	http.HandleFunc("/api/v1/averages/", averagesHandler)
	defer redisCli.Close()
	err := http.ListenAndServe(":8082", nil)
	log.Fatal(err)
}

// Get measures for an airport between two dates
func measuresHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 6 {

			// Get all the informations from the url
			aeroport := urlParts[4]
			nature := urlParts[5]
			queryValues := r.URL.Query()
			beginDate := queryValues.Get("beginDate")
			beginTime, err1 := time.Parse(layoutHeure, beginDate)
			endDate := queryValues.Get("endDate")
			endTime, err2 := time.Parse(layoutHeure, endDate)
			dateOk := beginTime.Before(endTime)

			if err1 == nil && err2 == nil && dateOk {

				var measuresStruct donneestruct.Measures
				var measures []donneestruct.Measure

				// Get data for each year
				for i := beginTime.Year(); i <= endTime.Year(); i++ {
					key := aeroport + ":" + nature + ":" + strconv.Itoa(i)

					res := getDataBetweenDates(key, beginTime, endTime)

					for _, value := range res {
						measure := donneestruct.Measure{CaptorID: value.CapteurID, Value: value.Valeur, Date: value.Date}
						measures = append(measures, measure)
					}
				}

				measuresStruct = donneestruct.Measures{Measures: measures}

				w.WriteHeader(http.StatusFound)
				writeJSON(w, measuresStruct)
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// Get the averages measures of an airport for one day
func averagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlParts := strings.Split(r.URL.Path, "/")
		if len(urlParts) == 5 {

			// Get all the informations from the url
			aeroport := urlParts[4]
			queryValues := r.URL.Query()
			date := queryValues.Get("date")
			dateTime, err := time.Parse(layoutDate, date)

			if err == nil {
				endDay := dateTime.Add(time.Hour * 24)
				endDay = endDay.Add(time.Second * -1)

				var averages donneestruct.Average
				var averageTemp float32
				var averagePres float32
				var averageWind float32

				// Get averages for each type of measure

				key := aeroport + ":" + donneestruct.TEMP + ":" + strconv.Itoa(dateTime.Year())
				resTemp := getDataBetweenDates(key, dateTime, endDay)

				for _, value := range resTemp {
					averageTemp = averageTemp + value.Valeur
				}

				if len(resTemp) != 0 {
					averageTemp = averageTemp / float32(len(resTemp))
				}

				key = aeroport + ":" + donneestruct.PRES + ":" + strconv.Itoa(dateTime.Year())
				resPres := getDataBetweenDates(key, dateTime, endDay)

				for _, value := range resPres {
					averagePres = averagePres + value.Valeur
				}

				if len(resPres) != 0 {
					averagePres = averagePres / float32(len(resPres))
				}

				key = aeroport + ":" + donneestruct.WIND + ":" + strconv.Itoa(dateTime.Year())
				resWind := getDataBetweenDates(key, dateTime, endDay)

				for _, value := range resWind {
					averageWind = averageWind + value.Valeur
				}

				if len(resWind) != 0 {
					averageWind = averageWind / float32(len(resWind))
				}

				averages = donneestruct.Average{AverageTemp: averageTemp, AveragePres: averagePres, AverageWind: averageWind}

				w.WriteHeader(http.StatusFound)
				writeJSON(w, averages)
				return
			}
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// Get data between two dates using a key
func getDataBetweenDates(key string, beginTime time.Time, endTime time.Time) []donneestruct.DonneesCapteur {
	// fmt.Println("key : ", key)
	var data []donneestruct.DonneesCapteur
	res, err := redis.Values(redisCli.Do("ZRANGEBYSCORE", key, beginTime.Unix(), endTime.Unix()))
	if err != nil {
		fmt.Println(err)
		return data
	}
	for len(res) > 0 {
		var message []byte
		res, err = redis.Scan(res, &message)
		if err != nil {
			fmt.Println(err)
		}

		var donneeCapt donneestruct.DonneesCapteur
		errJSON := json.Unmarshal(message, &donneeCapt)
		if errJSON != nil {
			fmt.Println(errJSON)
		}
		data = append(data, donneeCapt)
	}
	// fmt.Println("res : ", data)
	return data
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
