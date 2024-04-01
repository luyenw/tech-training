package main

import (
	"context"
	"fmt"
	"redis/config"
)

func main() {
	rdb := config.GetRedis()
	fmt.Println(rdb)
	val := rdb.Get(context.Background(), "connected")
	fmt.Println(val)

}
