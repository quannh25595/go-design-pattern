package creational

import "sync"

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func GetInstance() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

// Most of time, singleton reduces the testability
