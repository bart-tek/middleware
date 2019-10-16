package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

}

type person struct {
	id   int
	name string
}

var persons []person

func init() {
	persons = []person{
		person{id: 1, name: "JBey"},
		person{id: 2, name: "JBey"},
		person{id: 3, name: "JBey"},
		person{id: 4, name: "JBey"},
		person{id: 5, name: "JBey"},
		person{id: 6, name: "JBey"},
		person{id: 7, name: "JBey"},
		person{id: 8, name: "JBey"},
		person{id: 9, name: "JBey"},
		person{id: 10, name: "JBey"},
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

func personHandler(w http.ResponseWriter, r *http.Request) {
	strPath := r.URL.Path
	arrPath := strings.Split(strPath, "/")
	strID := arrPath[len(arrPath)-1]
	ID, err := strconv.Atoi(strID)
	if err != nil {
		log.Fatal(err)
	}
	if ID > 10 || ID < 1 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusFound)
		newPerson := person{}
		newPerson.name = "jbey"
		newPerson.id = ID
		writeJSON(w, newPerson)
	}
}

func main() {
	http.HandleFunc("/api/v1/person/", personHandler)
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}
