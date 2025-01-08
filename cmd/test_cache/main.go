package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"willnorris.com/go/imageproxy"
)

func main() {
	// Create a simple memory cache
	cache := make(map[string][]byte)
	testCache := &testCache{cache: cache}

	// Create the proxy with our cache
	proxy := imageproxy.NewProxy(http.DefaultTransport, testCache)
	
	// Set a long cache max age
	proxy.CacheMaxAge = 24 * time.Hour

	// Start the server
	fmt.Println("Starting server on :8080")
	fmt.Println("Try accessing: http://localhost:8080/300/https://octodex.github.com/images/codercat.jpg")
	fmt.Println("The first request will fetch from remote, subsequent requests will use cache")
	
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

// testCache is a simple in-memory cache implementation
type testCache struct {
	cache map[string][]byte
}

func (c *testCache) Get(key string) ([]byte, bool) {
	fmt.Printf("Cache lookup for key: %s\n", key)
	data, ok := c.cache[key]
	if ok {
		fmt.Println("Cache HIT")
	} else {
		fmt.Println("Cache MISS")
	}
	return data, ok
}

func (c *testCache) Set(key string, data []byte) {
	fmt.Printf("Caching data for key: %s (%d bytes)\n", key, len(data))
	c.cache[key] = data
}

func (c *testCache) Delete(key string) {
	delete(c.cache, key)
} 