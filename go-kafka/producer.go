package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
	"time"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}
type Order struct {
	OrderTime int64   `json:"ordertime"`
	OrderID   int     `json:"orderid"`
	ItemID    string  `json:"itemid"`
	Address   Address `json:"address"`
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "demo-topic"
	for i := 20; ; i++ {
		obj := Order{
			OrderTime: time.Now().Unix(),
			OrderID:   i,
			ItemID:    strconv.Itoa(i),
			Address: Address{
				City:    "Mountain View",
				State:   "CA",
				ZipCode: 94041,
			},
		}
		fmt.Printf("%+v\n", obj)

		bytes, err := json.Marshal(obj)
		if err != nil {
			fmt.Println(err)
			return
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          bytes,
		}, nil)

		time.Sleep(1 * time.Second)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
