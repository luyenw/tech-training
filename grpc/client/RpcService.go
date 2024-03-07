package client

import "context"

type IRpcService interface {
	Sum(a int32, b int32, ctx context.Context) (int32, error)
	Primes(num int32, ctx context.Context) (<-chan int, error)
	Avg(ch chan<- int32, ctx context.Context) (float32, error)
	Max(ch chan<- float32, ctx context.Context) (<-chan float32, error)
}
