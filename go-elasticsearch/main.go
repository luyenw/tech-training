package main

import (
	"encoding/json"
	"fmt"
	"go-es/config"
	"log"
	"strings"
)

type Hit struct {
	Index  string          `json:"_index"`
	Type   string          `json:"_type"`
	ID     string          `json:"_id"`
	Score  float64         `json:"_score"`
	Source json.RawMessage `json:"_source"`
}

type Response struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

func main() {
	query := `{
  "query": {
    "match": {
      "bytes": "1102"
    }
  }
}`
	res, err := config.GetES().Search(
		config.GetES().Search.WithIndex("kibana_sample_data_logs"),
		config.GetES().Search.WithBody(strings.NewReader(query)),
		config.GetES().Search.WithPretty(),
	)
	if err != nil {
		return
	}
	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	for _, hit := range response.Hits.Hits {
		var sourceMap map[string]interface{}
		err = json.Unmarshal(hit.Source, &sourceMap)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Index: %s\n Type: %s\n ID: %s\n Score: %f\n Source: %v\n", hit.Index, hit.Type, hit.ID, hit.Score, sourceMap)
	}
}
