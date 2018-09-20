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
	f, _ := os.Open("tmp")
	s := bufio.NewScanner(f)
	var data []string
	for s.Scan() {
		data = append(data, s.Text())
	}
	data = data[10:]

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "censoredplanet-v1")
	if err != nil {
		panic(err)
	}
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

	it, err := q.Read(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Println(values[0].(string))
	}
}
