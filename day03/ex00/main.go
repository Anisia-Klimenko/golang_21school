package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Places struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

type GeoPoint struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}

func main() {
	log.SetFlags(0)

	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},

		// A simple incremental backoff function
		//
		RetryBackoff: func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },

		// Retry up to 5 attempts
		//
		MaxRetries: 5,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Create the indexer
	//
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     es,     // The Elasticsearch client
		Index:      "test", // The default index name
		NumWorkers: 4,      // The number of worker goroutines (default: number of CPUs)
		FlushBytes: 5e+6,   // The flush threshold in bytes (default: 5M)
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	// Add an item to the indexer
	//
	err = indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			// Action field configures the operation to perform (index, create, delete, update)
			Action: "index",

			// DocumentID is the optional document ID
			DocumentID: "1",

			// Body is an `io.Reader` with the payload
			Body: strings.NewReader(`{"title":"Test"}`),

			// OnSuccess is the optional callback for each successful operation
			OnSuccess: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem,
			) {
				fmt.Printf("[%d] %s test/%s", res.Status, res.Result, item.DocumentID)
			},

			// OnFailure is the optional callback for each failed operation
			OnFailure: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem, err error,
			) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		},
	)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := indexer.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	//
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		log.Fatalf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
	} else {
		log.Printf("Successfully indexed [%d] documents", stats.NumFlushed)
	}
}

func main2() {
	var (
	//countSuccessful uint64
	)

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	index := "places"

	schema, err := os.ReadFile("schema.json")

	mapping := `
	{
	 "settings": {
	   "number_of_shards": 1
	 },
	 "mappings": ` + string(schema) + `
	}`

	res, err := client.Indices.Create(
		index,
		client.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	////// JSON //////

	//f, err := os.Open("materials/data.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//
	//csvReader := csv.NewReader(f)
	//csvReader.Comma = '\t'
	//csvReader.FieldsPerRecord = -1
	//data, err := csvReader.ReadAll()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//places := createPlaces(data)
	//jsonData, err := json.MarshalIndent(places, "", "  ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(jsonData))

	////// JSON //////

	//bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
	//	Index:  index,  // Имя индекса по умолчанию
	//	Client: client, // Elasticsearch клиент
	//})

	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})
	indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			Action: "place",
			Body:   strings.NewReader(`{"place":"0"}`),
		})
	indexer.Close(context.Background())

	//for _, p := range places {
	//	data, err := json.Marshal(p)
	//	if err != nil {
	//		log.Fatalf("Cannot encode place: %s", err)
	//	}
	//	err = bi.Add(
	//		context.Background(),
	//		esutil.BulkIndexerItem{
	//			Action: index + "/place",
	//			Body:   bytes.NewReader(data),
	//			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
	//				atomic.AddUint64(&countSuccessful, 1)
	//			},
	//			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
	//				if err != nil {
	//					log.Printf("ERROR: %s", err)
	//				} else {
	//					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
	//				}
	//			},
	//		})
	//	if err != nil {
	//		log.Fatalf("Unexpected error: %s", err)
	//	}
	//}
	//if err := bi.Close(context.Background()); err != nil {
	//	log.Fatalf("Unexpected error: %s", err)
	//}
	//biStats := bi.Stats()
	//if biStats.NumFailed > 0 {
	//	log.Fatalf("Error")
	//} else {
	//	log.Printf("Success")
	//}

	//////////////////////////

	//req, err := http.NewRequest("PUT", "http://localhost:9200/_mapping?include_type_name=true", bytes.NewBuffer(jsonData))
	//req.Header.Set("Content-Type", "application/x-ndjson")
	//if err != nil {
	//	log.Fatalf("http.NewRequest ERROR:", err)
	//} else {
	//	fmt.Println("HTTP Request:", req)
	//}
	//clientHTTP := &http.Client{}
	//
	//// Pass HTTP request to Elasticsearch client and check for errors
	//resp, err := clientHTTP.Do(req)
	//if err != nil {
	//	log.Fatalf("client.Do() ERROR:", err)
	//}
	//
	//// Close the response body after operations are complete
	//defer resp.Body.Close()

}

type BulkIndexerItem struct {
	Index           string
	Action          string
	DocumentID      string
	Body            io.Reader
	RetryOnConflict *int

	OnSuccess func(context.Context, BulkIndexerItem, esutil.BulkIndexerResponseItem)        // Для каждого элемента
	OnFailure func(context.Context, BulkIndexerItem, esutil.BulkIndexerResponseItem, error) // Для каждого элемента
}

func createPlaces(data [][]string) []Places {
	var places []Places

	for i, line := range data {
		if i > 0 {
			var rec Places
			for j, field := range line {
				if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Address = field
				} else if j == 3 {
					rec.Phone = field
				} else if j == 4 {
					rec.Location.Lon = field
				} else if j == 5 {
					rec.Location.Lat = field
				}
			}
			places = append(places, rec)
		}
	}
	return places
}
