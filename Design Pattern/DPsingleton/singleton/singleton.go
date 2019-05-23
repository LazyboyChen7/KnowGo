package singleton

import (
	"sync"
)

type king struct {
	Name    string
	Age     int
	Country string
	Extra   byte
	Final   byte
}

var (
	once     sync.Once
	instance *king
)

func New(name string, age int, country string) *king {
	once.Do(func() {
		instance = &king{
			Name:    name,
			Age:     age,
			Country: country,
			Extra:   '1',
			Final:   '1',
		}
	})
	return instance
}
