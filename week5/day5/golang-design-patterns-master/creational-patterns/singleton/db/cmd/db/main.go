package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/creational-patterns/singleton/db"
)

func main() {
	value, err := db.Repository().Get("id")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
