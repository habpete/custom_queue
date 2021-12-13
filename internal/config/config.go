package config

type Config interface {
	Get() interface{}
	Set()
}

func New() Config {
	return &config{}
}

type config struct{}

func (i *config) Get() interface{} {
	return nil
}

func (i *config) Set() {}
