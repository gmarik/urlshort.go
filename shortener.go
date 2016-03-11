package urlshort

import (
	"fmt"
	"hash"
	"sync"

	"github.com/spaolacci/murmur3"
)

type Hash uint32

type URLShortener interface {
	ShortUrl(string) string
}

type HashShortener struct {
	// TODO: user hasher Pool?
	mu     *sync.Mutex
	hasher hash.Hash32

	baseDomain string
}

func NewHashShortener(domain string) *HashShortener {
	return &HashShortener{
		mu:         &sync.Mutex{},
		hasher:     murmur3.New32(),
		baseDomain: domain,
	}
}

func (s *HashShortener) Hash(val string) Hash {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.hasher.Reset()
	s.hasher.Write([]byte(val))

	return Hash(s.hasher.Sum32())
}

func (s *HashShortener) ShortUrl(url string) string {
	return fmt.Sprintf("%s/%x", s.baseDomain, s.Hash(url))
}
