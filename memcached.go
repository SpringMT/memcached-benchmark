package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedClient struct {
	Client *memcache.Client
}

func (mc *MemcachedClient) increment(key string) (uint64, error) {
	incr, err := mc.Client.Increment(key, 1)
	return incr, err
}

func (mc *MemcachedClient) close() {
}

// MemcachedNewClient Redisクライアントを返す
// https://godoc.org/github.com/go-redis/redis#Options
func MemcachedNewClient(host string) MemcachedClient {
	client := memcache.New(host + ":11211")
	return MemcachedClient{Client: client}
}
