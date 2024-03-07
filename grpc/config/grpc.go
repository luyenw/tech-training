package config

import (
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
	"sync"
)

var (
	mu     sync.Mutex
	client proto.CalculatorServiceClient
)

func GetRpcClient() proto.CalculatorServiceClient {
	if client == nil {
		mu.Lock()
		defer mu.Unlock()
		cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())
		if err != nil {
			log.Fatalf(err.Error())
		}
		client = proto.NewCalculatorServiceClient(cc)
	}
	return client
}
