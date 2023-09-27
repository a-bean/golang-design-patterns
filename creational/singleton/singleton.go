package singleton

import "sync"

// 饿汉式
type Singleton struct{}

var instance *Singleton

func init() {
	instance = &Singleton{}
}

func GetInstance() *Singleton {
	return instance
}

// 懒汉式（双重检测）
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
