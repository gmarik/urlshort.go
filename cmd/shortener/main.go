package main

import (
	"log"
	"net/http"

	"github.com/gmarik/urlshort.go"
)

func main() {

	service := &urlshort.Service{
		Store:     urlshort.NewMapStore(),
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
