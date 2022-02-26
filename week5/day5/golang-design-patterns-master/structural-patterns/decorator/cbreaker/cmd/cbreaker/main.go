package main

import (
	"fmt"
	"time"

	"github.com/svett/golang-design-patterns/structural-patterns/decorator/cbreaker"
)

func main() {
	repository := &cbreaker.Repository{}
	retrier := &cbreaker.Retrier{
		RetryCount:   5,
		WaitInterval: time.Second,
		Fetcher:      repository,
	}

	data, err := repository.Fetch(cbreaker.Args{"id": "1"})
	fmt.Printf("#1 repository.Fetch: %v\n", data)

	data, err = retrier.Fetch(cbreaker.Args{})
	fmt.Printf("#2 retrier.Fetch error: %v\n", err)

	data, err = retrier.Fetch(cbreaker.Args{"id": "1"})
	fmt.Printf("#3 retrier.Fetch: %v\n", data)
}
