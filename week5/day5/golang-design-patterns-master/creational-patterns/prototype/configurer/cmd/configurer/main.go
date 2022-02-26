package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/creational-patterns/prototype/configurer"
)

func main() {
	config := configurer.NewConfig("guest", "/home/guest")
	rootConfig := config.WithUser("root").WithWorkDir("/root")

	fmt.Println("Guest Config", config)
	fmt.Println("Root Config", rootConfig)
}
