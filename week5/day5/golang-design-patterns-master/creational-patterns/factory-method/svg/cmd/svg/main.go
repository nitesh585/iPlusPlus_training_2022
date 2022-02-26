package main

import (
	"os"

	"github.com/svett/golang-design-patterns/creational-patterns/factory-method/svg"
)

func main() {
	doc := &svg.Document{
		ShapeFactories: []svg.ShapeFactory{
			&svg.CircleFactory{},
			&svg.RactangleFactory{},
		},
	}

	doc.Draw(os.Stdout)
}
