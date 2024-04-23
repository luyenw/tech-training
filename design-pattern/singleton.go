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
	if object == nil {
		mu.Lock()
		object = &Object{createdAt: time.Now().String()}
		defer mu.Unlock()
	}
	return object
}

func main() {
	for i := 0; i < 10; i++ {
		obj := getInstance()
		fmt.Printf("%+v\n", obj)
	}
}
