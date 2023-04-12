package cache

import (
	"errors"
	"time"
)

type Value struct {
	Value interface{}
	Ttl   time.Time
}

type dictionary map[string]Value

type Cache struct {
	cacheMemory dictionary
}

func New() Cache {
	return Cache{
		cacheMemory: dictionary{},
	}
}

func (c Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.cacheMemory[key] = Value{value, time.Now().Add(ttl)}
}

func (c Cache) Delete(key string) {
	_, contains := c.cacheMemory[key]

	if !contains {
		panic("Can't delete not existing key!")
	}

	delete(c.cacheMemory, key)
}

func (c Cache) Get(key string) (interface{}, error) {
	value, contains := c.cacheMemory[key]

	if contains && value.Ttl.Second() < time.Now().Second() {
		c.Delete(key)
		contains = false
	}

	if contains {
		return value, nil
	}

	return "Such key doesn't exist yet!", errors.New("Such key doesn't exist yet!")
}
