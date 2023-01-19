# go-nofluffjobs

go-nofluffjobs is a Go client library for searching job offers on European IT job board - No Fluff Jobs.

# Installation
Command line:
```
go get github.com/necsord/go-nofluffjobs
```
# Usage
```go
package main

import (
	"context"
	"github.com/necsord/go-nofluffjobs/nofluffjobs"
	"log"
)

func main() {
	// Create client
	client, err := nofluffjobs.NewClient("https://nofluffjobs.com/api", nil, nil)
	if err != nil {
		log.Fatalf("could not create nofluffjobs client: %v", err)
	}

	// Prepare query params and request body
	query := nofluffjobs.SearchPostingQuery{
		Limit:          50,
		Offset:         0,
		SalaryCurrency: "PLN",
		SalaryPeriod:   "hour",
		Region:         "pl",
	}
	body := nofluffjobs.SearchPostingRequest{
		Page: 1,
		CriteriaSearch: nofluffjobs.CriteriaSearch{
			Requirement: []string{"Golang"},
			Employment:  []string{"b2b"},
		},
	}
	// Make request
	postings, err := client.SearchPosting(context.Background(), query, body)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	// Log response
	log.Println(postings)
}
```