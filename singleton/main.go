package main

import "sync"

var once sync.Once
var instance *singletonDatabase

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func main() {

}

func GetSingoletonDatabase() *singletonDatabase {
	once.Do(func() {
		db := singletonDatabase{}
		//read
		instance = &db
	})
	return instance
}
