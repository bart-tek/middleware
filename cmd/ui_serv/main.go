package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<head></head>")
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, "<h1>Hello Evrard-Nil !</h1>")
	fmt.Fprintf(w, "<p>")
	fmt.Fprintf(w, "path:"+r.URL.Path+"</p>")
	fmt.Fprintf(w, "host:"+r.URL.Host+"</p>")
	fmt.Fprintf(w, "uri:"+r.URL.RequestURI()+"</p>")
	fmt.Fprintf(w, "query:"+r.URL.RawQuery+"</p>")
	fmt.Fprintf(w, "</p>")
	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "D:/Profiles/edaillet/go/src/github.com/evrard-nil/middleware/cmd/ui_serv/www/main.html")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}
