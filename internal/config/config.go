package config

import "sync"

type Config interface {
	Get(key string) *Value
	Set(key string, value interface{})
}

func New() Config {
	return &config{
		values: make(map[string]*Value),
	}
}

type config struct {
	valuesMtx sync.RWMutex
	values    map[string]*Value
}

func (i *config) Get(key string) *Value {
	i.valuesMtx.RLock()
	defer i.valuesMtx.RUnlock()

	return i.values[key]
}

func (i *config) Set(key string, value interface{}) {
	i.valuesMtx.Lock()
	i.values[key] = NewValue(value)
	i.valuesMtx.Unlock()
}
