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
	mu.Lock()
	defer mu.Unlock()
	if client == nil {
		cc, err := grpc.Dial("localhost:50069", grpc.WithInsecure())
		if err != nil {
			log.Fatalf(err.Error())
		}
		client = proto.NewCalculatorServiceClient(cc)
	}
	return client
}
