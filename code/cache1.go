package main

type Cache struct {
	m map[int]int
}

func (c *Cache) Set(k int, v int) {
	m[k] = v
}

func (c *Cache) Get(k int) {
	return m[k]
}
