package main

import (
	"bufio"
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"os"
)

func main() {
	// Load a file containing IPs for parameterized query argument
	f, err := os.Open("tmp")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var data []string
	for s.Scan() {
		data = append(data, s.Text())
	}

	// Create context and client for BigQuery
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "censoredplanet-v1")
	if err != nil {
		panic(err)
	}

	// Create SQL query with data slice as argument
	q := client.Query(`
		SELECT * FROM satellite.answers
		WHERE
		answer IN UNNEST(@data)
	`)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "data",
			Value: data,
		},
	}

	// Execute query
	it, err := q.Read(ctx)
	if err != nil {
		panic(err)
	}

	// Loop through results, print to stdout
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(values[0].(string))
	}
}
