package urlshort

import (
	"testing"
)

func TestShortenerHash(t *testing.T) {

	s := NewHashShortener("http://u.ly")

	tcases := []struct {
		val      string
		hash     Hash
		shortUrl string
	}{
		{"", 0x00000000, "http://u.ly/0"},
		{"http://gmarik.info", 0x5595bafb, "http://u.ly/5595bafb"},
	}

	for _, tcase := range tcases {

		if got := s.Hash(tcase.val); got != tcase.hash {
			t.Errorf("'%s': 0x%08x (want 0x%x)", tcase.val, got, tcase.hash)
		}
		if got := s.ShortUrl(tcase.val); got != tcase.shortUrl {
			t.Errorf("'%s': %s (want %s)", tcase.val, got, tcase.shortUrl)
		}
	}
}
