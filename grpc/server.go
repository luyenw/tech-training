package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
)

type Server struct {
	proto.CalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	// implementation
	result := in.GetFirst() + in.GetSecond()
	//
	return &proto.SumResponse{Result: result}, nil
}

func (s *Server) Primes(in *proto.PrimesRequest, stream proto.CalculatorService_PrimesServer) error {
	//
	num := in.GetNumber()
	var k int32
	var err error
	k = 2
	for num > 1 {
		if num%k == 0 {
			num /= k
			time.Sleep(1 * time.Second)
			err = stream.Send(&proto.PrimesResponse{
				Result: k,
			})
			if err != nil {
				return err
			}
		} else {
			k += 1
		}
	}
	//
	return nil
}

func (s *Server) Avg(stream proto.CalculatorService_AvgServer) error {
	sum := 0
	n := 0
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(recv.GetNumber())
		sum += int(recv.GetNumber())
		n += 1
		if err != nil {
			return err
		}
	}
	avg := float32(sum) / float32(n)
	fmt.Println(avg)
	return stream.SendAndClose(&proto.AvgResponse{Result: avg})
}

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	_max := -1.0
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_max = math.Max(_max, float64(recv.GetNumber()))
		err = stream.Send(&proto.MaxResponse{Result: float32(_max)})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf(err.Error())
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
