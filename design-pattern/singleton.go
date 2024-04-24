package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	createdAt string
}

var (
	object *Object
	mu     sync.Mutex
)

func getInstance() *Object {
	mu.Lock()
	defer mu.Unlock()
	if object == nil {
		object = &Object{createdAt: time.Now().String()}
	}
	return object
}

func main() {
	for i := 0; i < 10; i++ {
		obj := getInstance()
		fmt.Printf("%+v\n", obj)
	}
}
