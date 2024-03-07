package client

import (
	"context"
	"grpc/config"
	"grpc/proto"
	"io"
	"log"
)

type RpcService struct{}

func (srv RpcService) Sum(a int32, b int32, ctx context.Context) (int32, error) {
	client := config.GetRpcClient()
	sum, err := client.Sum(ctx, &proto.SumRequest{First: a, Second: b})
	if err != nil {
		return 0, err
	}
	return sum.GetResult(), nil
}

func (srv RpcService) Primes(ch chan int32, num int32, ctx context.Context) error {
	client := config.GetRpcClient()
	primes, err := client.Primes(ctx, &proto.PrimesRequest{Number: num})
	if err != nil {
		return err
	}

	for {
		k, err := primes.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf(err.Error())
		}
		res := k.GetResult()
		ch <- res
	}
	close(ch)
	return nil
}

func (srv RpcService) Avg(ch chan int32, ctx context.Context) (float32, error) {
	client := config.GetRpcClient()
	avg, err := client.Avg(ctx)
	if err != nil {
		return 0, err
	}
	var n int32
	opened := true
	for opened {
		n, opened = <-ch
		if opened {
			err := avg.Send(&proto.AvgRequest{Number: n})
			if err != nil {
				return 0, err
			}
		}
	}
	recv, err := avg.CloseAndRecv()
	if err != nil {
		return 0, err
	}
	return recv.GetResult(), nil
}
