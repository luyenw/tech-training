package main

import (
	"context"
	"fmt"
	"grpc/client"
	"grpc/config"
	"grpc/proto"
	"io"
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
	c := config.GetRpcClient()
	maxClient, err := c.Max(context.Background())
	if err != nil {
		return
	}
	ch2 := make(chan int)
	go func() {
		li := [6]float32{1, 5, 3, 6, 2, 20}
		for i := 0; i < len(li); i++ {
			fmt.Printf("sent: %.1f\n", li[i])
			time.Sleep(1 * time.Second)
			err := maxClient.Send(&proto.MaxRequest{Number: li[i]})
			if err != nil {
				log.Fatalf(err.Error())
				return
			}
		}
		err := maxClient.CloseSend()
		if err != nil {
			return
		}
	}()
	go func() {
		for {
			recv, err := maxClient.Recv()
			if err == io.EOF {
				ch2 <- 1
			}
			if err != nil {
				return
			}
			fmt.Printf("received: %.1f\n", recv.GetResult())
		}
	}()
	<-ch2
}
