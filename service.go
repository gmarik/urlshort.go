package urlshort

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Shorturl string `json:"shorturl"`
	Longurl  string `json:"longurl"`
}

type Service struct {
	Store     Store
	Shortener URLShortener
}

func NewService(baseUrl string) *Service {
	return &Service{
		Store:     NewMapStore(),
		Shortener: NewHashShortener(baseUrl),
	}
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	longurl := r.FormValue("longurl")
	shorturl := s.Shortener.ShortUrl(longurl)
	err := s.Store.Put(shorturl, longurl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	resp := &Response{
		Shorturl: shorturl,
		Longurl:  longurl,
	}
	s.write(w, resp)
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	shorturl := r.FormValue("shorturl")
	longurl, err := s.Store.Get(shorturl)
	if err != nil {
		if err == NotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	resp := &Response{
		Shorturl: shorturl,
		Longurl:  longurl,
	}

	s.write(w, resp)
}

func (s *Service) write(w http.ResponseWriter, resp *Response) {
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
}
