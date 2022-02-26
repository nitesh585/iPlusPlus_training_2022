package cbreaker

import (
	"fmt"
	"time"
)

// Args of fetching function
type Args map[string]string

// Data returned by fetch
type Data map[string]string

// Fetcher fetches a data from remote endpoint
type Fetcher interface {
	// Fetch fetches the data
	Fetch(args Args) (Data, error)
}

// Retrier retries multiple times
type Retrier struct {
	RetryCount   int
	WaitInterval time.Duration
	Fetcher      Fetcher
}

// Fetch fetches data
func (r *Retrier) Fetch(args Args) (Data, error) {
	for retry := 1; retry <= r.RetryCount; retry++ {
		fmt.Printf("Retrier retries to fetch for %d\n", retry)
		if data, err := r.Fetcher.Fetch(args); err == nil {
			fmt.Printf("Retrier fetched for %d\n", retry)
			return data, nil
		} else if retry == r.RetryCount {
			fmt.Printf("Retrier failed to fetch for %d times\n", retry)
			return Data{}, err
		}
		fmt.Printf("Retrier is waiting after error fetch for %v\n", r.WaitInterval)
		time.Sleep(r.WaitInterval)
	}

	return Data{}, nil
}

// Repository of data
type Repository struct{}

// Fetch fetches data
func (r *Repository) Fetch(args Args) (Data, error) {
	if len(args) == 0 {
		return Data{}, fmt.Errorf("No arguments are provided")
	}

	data := Data{
		"user":     "root",
		"password": "swordfish",
	}
	fmt.Printf("Repository fetched data successfully: %v\n", data)
	return data, nil
}
