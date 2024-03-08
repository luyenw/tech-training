package main

import (
	"context"
	"fmt"
	"grpc/client"
	"log"
	"time"
)

func main() {
	var srv client.RpcService
	//
	log.Println("unary")
	sum, err := srv.Sum(3, 5, context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(sum)
	//
	log.Println("server streaming")
	ch := make(chan int32)
	fmt.Println("sent: 120")
	go func() {
		err = srv.Primes(ch, 120, context.Background())
	}()
	opened := true
	var k int32
	fmt.Print("received: ")
	for opened {
		k, opened = <-ch
		if opened {
			fmt.Printf("%d, ", k)
		} else {
			fmt.Print("\n")
		}
	}
	//
	log.Println("client streaming")
	ch1 := make(chan int32)
	fmt.Print("sent: ")
	go func() {
		for i := 1; i <= 4; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- int32(i)
			fmt.Printf("%d, ", i)
		}
		close(ch1)
	}()
	avg, err := srv.Avg(ch1, context.Background())
	if err != nil {
		return
	}
	fmt.Printf("\nreceived: %f\n", avg)

	// bi-directional streaming
	log.Println("bi-directional streaming")
	in := make(chan float32)
	go func() {
		li := []float32{1, 5, 3, 6, 2, 20}
		for _, i := range li {
			time.Sleep(1 * time.Second)
			in <- i
			fmt.Printf("sent: %.1f\n", i)
		}
		close(in)
	}()
	out, err := srv.Max(in, context.Background())
	opened = true
	var re float32
	for opened {
		re, opened = <-out
		if opened {
			fmt.Printf("received: %.1f\n", re)
		}
	}
}
