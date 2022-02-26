package svg

import (
	"fmt"
	"io"
	"math"
)

type Shape interface {
	Draw(io.Writer) error
}

type Point struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Viewport struct {
	Location Point
	Size     Size
}

type Document struct {
	ShapeFactories []ShapeFactory
}

func (doc *Document) Draw(w io.Writer) error {
	viewport := Viewport{
		Location: Point{
			X: 0,
			Y: 0,
		},
		Size: Size{
			Width:  640,
			Height: 480,
		},
	}
	if _, err := fmt.Fprintf(w, `<svg height="%f" width="%f">`, viewport.Size.Height, viewport.Size.Width); err != nil {
		return err
	}

	for _, factory := range doc.ShapeFactories {
		shape := factory.Create(viewport)
		if err := shape.Draw(w); err != nil {
			return err
		}
	}

	_, err := fmt.Fprint(w, `</svg>`)
	return err
}

type ShapeFactory interface {
	Create(viewport Viewport) Shape
}

type Circle struct {
	Location Point
	Radius   float64
}

func (c *Circle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<circle cx="%f" cy="%f" r="%f"/>`, c.Location.X, c.Location.Y, c.Radius)
	return err
}

type CircleFactory struct{}

func (factory *CircleFactory) Create(viewport Viewport) Shape {
	return &Circle{
		Location: viewport.Location,
		Radius:   math.Min(viewport.Size.Width, viewport.Size.Height),
	}
}

type Rectangle struct {
	Location Point
	Size     Size
}

func (rect *Rectangle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<rect x="%f" y="%f" width="%f" height="%f"/>`, rect.Location.X, rect.Location.Y, rect.Size.Width, rect.Size.Height)
	return err
}

type RactangleFactory struct{}

func (factory *RactangleFactory) Create(viewport Viewport) Shape {
	return &Rectangle{
		Location: viewport.Location,
		Size:     viewport.Size,
	}
}
