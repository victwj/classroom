package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "censoredplanet-v1")
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
}
