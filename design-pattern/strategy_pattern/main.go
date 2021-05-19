package main

import "fmt"

// https://refactoringguru.cn/design-patterns/strategy/go/example

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct {
}

func (f *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo algo.")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru algo.")
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lru := &lru{}
	cache := initCache(lru)

	cache.add("a", "1")
	cache.add("b", "2")

	// evict
	cache.add("c", "3")

	cache.setEvictionAlgo(&fifo{})

	// evict
	cache.add("d", "4")
}
