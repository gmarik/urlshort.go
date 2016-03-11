package urlshort

import (
	"fmt"
	"sync"
)

var NotFound error = fmt.Errorf("Not Found")

type Store interface {
	Get(string) (string, error)
	Put(string, string) error
}

func NewMapStore() *MapStore {
	return &MapStore{
		mu: &sync.Mutex{},
		m:  make(map[string]string),
	}
}

type MapStore struct {
	m  map[string]string
	mu *sync.Mutex
}

func (ms MapStore) Get(val string) (string, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	v, ok := ms.m[val]
	if !ok {
		return "", NotFound
	}

	return v, nil
}
func (ms MapStore) Put(key, val string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.m[key] = val
	return nil
}
