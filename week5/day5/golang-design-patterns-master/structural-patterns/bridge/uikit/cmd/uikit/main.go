package main

import (
	"fmt"

	"github.com/svett/golang-design-patterns/structural-patterns/bridge/uikit"
)

func main() {
	openGL := &uikit.OpenGL{}
	direct2D := &uikit.Direct2D{}

	circle := &uikit.Circle{
		Center: uikit.Point{X: 100, Y: 100},
		Radius: 50,
	}

	fmt.Println()

	circle.DrawingContext = openGL
	circle.Draw()

	fmt.Println()

	circle.DrawingContext = direct2D
	circle.Draw()
}
