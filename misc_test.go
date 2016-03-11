package urlshort

import (
	"hash"
	"testing"

	"github.com/spaolacci/murmur3"
)

func TestMurmurHash(t *testing.T) {

	var (
		h32 hash.Hash32 = murmur3.New32()
		in              = "Hello, world"

		exp uint32 = 0x6a728c54
	)

	h32.Write([]byte(in))

	if got := h32.Sum32(); exp != got {
		t.Errorf("'%s': 0x%x (want 0x%x)", in, got, exp)
	}
}
