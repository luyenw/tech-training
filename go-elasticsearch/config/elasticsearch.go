package config

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"sync"
	"time"
)

var (
	instance *elasticsearch.Client
	err      error
	mutex    sync.Mutex
)

func initES() *elasticsearch.Client {
	var es *elasticsearch.Client
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	_, err = es.Info(es.Info.WithContext(ctx))
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	return es
}
func GetES() *elasticsearch.Client {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		instance = initES()
		fmt.Println(time.Now())
	}
	return instance
}
