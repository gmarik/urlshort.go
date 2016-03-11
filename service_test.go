package urlshort

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService(t *testing.T) {

	service := &Service{
		Store:     NewMapStore(),
		Shortener: NewHashShortener("http://u.ly"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			service.Get(w, r)
		case "POST":
			service.Create(w, r)
		}
	}))
	defer ts.Close()

	longUrl := "http://gmarik.info"

	_, err := http.Post(ts.URL+"?longurl="+longUrl, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	shortUrl := service.Shortener.ShortUrl(longUrl)

	res, err := http.Get(ts.URL + "?shorturl=" + shortUrl)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("Unexpected Response:", string(body))
	}

	var resp Response

	if err := json.Unmarshal(body, &resp); err != nil {
		t.Errorf("Error: %s, Body: %q", err, string(body))
	}

	if resp.Longurl != longUrl {
		t.Errorf("\nExp %s\nGot: %s", longUrl, resp.Longurl)
	}
	if resp.Shorturl != shortUrl {
		t.Errorf("\nExp %s\nGot: %s", shortUrl, resp.Shorturl)
	}
}
