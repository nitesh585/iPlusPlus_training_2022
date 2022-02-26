package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/creational-patterns/builder/messenger"
)

func main() {
	sender := &messenger.Sender{}

	jsonMsg, err := sender.BuildMessage(&messenger.JSONMessageBuilder{})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonMsg.Body))

	xmlMsg, err := sender.BuildMessage(&messenger.XMLMessageBuilder{})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xmlMsg.Body))
}
