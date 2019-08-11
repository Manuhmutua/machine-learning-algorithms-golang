package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds
	c := cache.New(5*time.Minute, 30*time.Second)

	// Put a key and value into the cache.
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}

}