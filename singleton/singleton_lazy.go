package singleton

import "sync"

type lazySingleton struct{}

var (
	lazyInstance *lazySingleton
	once         = &sync.Once{}
)

func GetInstanceLazy() *lazySingleton {
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &lazySingleton{}
		})
	}
	return lazyInstance
}
