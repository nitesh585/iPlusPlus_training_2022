package db

import (
	"fmt"
	"sync"
)

type Map interface {
	Set(key, data string)
	Get(key string) (string, error)
}

type repository struct {
	items map[string]string
	mu    sync.RWMutex
}

func (r *repository) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = data
}

func (r *repository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

var (
	r    *repository
	once sync.Once
)

func Get(key string) (string, error) {
	return r.Get(key)
}

func Repository() Map {
	if r == nil {
		once.Do(func() {
			r = &repository{
				items: make(map[string]string),
			}
		})
	}
	return r
}
