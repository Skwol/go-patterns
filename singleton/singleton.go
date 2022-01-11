package main

import (
	"fmt"
	"sync"
)

// easier to implement with sync.Once DO
var lock = &sync.Mutex{}

type singleton struct{}

var singletonIns *singleton

func getInstance() *singleton {
	if singletonIns == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonIns == nil {
			fmt.Println("creating singleton.")
			singletonIns = &singleton{}
		} else {
			fmt.Println("singleton already exist")
		}
	} else {
		fmt.Println("singlton already created")
	}

	return singletonIns
}
