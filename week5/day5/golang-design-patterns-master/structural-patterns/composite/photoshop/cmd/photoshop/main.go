package main

import "github.com/svett/golang-design-patterns/structural-patterns/composite/photoshop"

func main() {
	circle := &photoshop.Circle{
		Center: photoshop.Point{X: 100, Y: 100},
		Radius: 50,
	}

	square := &photoshop.Square{
		Location: photoshop.Point{X: 50, Y: 50},
		Side:     20,
	}

	layer := &photoshop.Layer{
		Elements: []photoshop.VisualElement{
			circle,
			square,
		},
	}

	layer.Draw(&photoshop.Drawer{})
}
