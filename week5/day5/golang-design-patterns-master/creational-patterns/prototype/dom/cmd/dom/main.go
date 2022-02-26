package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/creational-patterns/prototype/dom"
)

func main() {
	directorNode := dom.NewElement("Director of Engineering")

	engManagerNode := dom.NewElement("Engineering Manager")
	engManagerNode.AddChild(dom.NewElement("Lead Software Engineer"))

	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(engManagerNode)

	officeManagerNode := dom.NewElement("Office Manager")
	directorNode.AddChild(officeManagerNode)

	fmt.Println("")
	fmt.Println("# Company Hierarchy")
	fmt.Print(directorNode)
	fmt.Println("")
	fmt.Println("# Team Hiearachy")
	fmt.Print(engManagerNode.Clone())
}
