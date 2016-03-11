package urlshort

import (
	"testing"
)

func BenchmarkMapStore(t *testing.B) {
	s := NewHashShortener("http://u.ly")
	store := NewMapStore()

	val := "http://gmarik.info"
	key := s.ShortUrl(val)

	for i := 0; i < t.N; i += 1 {

		if err := store.Put(key, val); err != nil {
			t.Error(err)
		}
		if v, err := store.Get(key); err != nil || v != val {
			t.Errorf("\nGot: %s\nExp: %s", v, val)
		}
	}
}
