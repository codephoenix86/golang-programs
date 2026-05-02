package main

import "fmt"

type Storage interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}

type Cache struct {
	store map[string]any
}

func (cache *Cache) Set(key string, value any) {
	cache.store[key] = value
}

func (cache *Cache) Get(key string) (any, bool) {
	value, ok := cache.store[key]
	return value, ok
}

func newCache() *Cache {
	return &Cache{store: make(map[string]any)}
}

func main() {
	cache := newCache()
	cache.Set("name", "Naresh Lohar")
	cache.Set("age", 22)

	if name, found := cache.Get("name"); found {
		fmt.Println("Found user:", name)
	}

	if _, found := cache.Get("email"); !found {
		fmt.Println("Email not found in cache.")
	}
}
