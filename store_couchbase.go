package urlshort

import (
	"fmt"

	"github.com/couchbase/go-couchbase"
)

type Couchbase struct {
	b *couchbase.Bucket
}

func (c *Couchbase) Connect(url, bucket string) error {
	client, err := couchbase.Connect(url)
	if err != nil {
		return fmt.Errorf("Error connecting:  %v", err)
	}

	pool, err := client.GetPool("default")
	if err != nil {
		return fmt.Errorf("Error getting pool:  %v", err)
	}

	b, err := pool.GetBucket(bucket)
	if err != nil {
		fmt.Errorf("Error getting bucket:  %v", err)
	}

	c.b = b

	return nil
}

func (c *Couchbase) Put(shortUrl, longUrl string) error {
	return c.b.Set(shortUrl, 0, longUrl)
}

func (c *Couchbase) Get(shortUrl string) (string, error) {
	var s string
	if err := c.b.Get(shortUrl, &s); err != nil {
		return "", err
	}
	return s, nil
}
