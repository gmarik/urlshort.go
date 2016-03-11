package main

import (
	"log"
	"net/http"

	"github.com/gmarik/urlshort.go"
)

func main() {

	cb := &urlshort.Couchbase{}
	err := cb.Connect("couchbase://192.168.61.101:8091", "testload")
	if err != nil {
		log.Fatal(err)
	}

	service := &urlshort.Service{
		Store:     cb,
		Shortener: urlshort.NewHashShortener("http://u.ly"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			service.Get(w, r)
		case "POST":
			service.Create(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(":4040", nil))
}
