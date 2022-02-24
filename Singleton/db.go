package Singleton

import "sync"

var instance *db

type db struct {
}

func (d db) getUser(name string) string {
	return name
}

func getInstance() *db {
	if instance == nil {
		instance = &db{}
	}

	return instance
}

var lock sync.Mutex

func getInstanceWithLock() *db {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		instance = &db{}
	}
	return instance
}

func onceInstance() {
	instance = &db{}
}

var once sync.Once

func getInstanceWithOnce() *db {
	once.Do(onceInstance)

	return instance
}
