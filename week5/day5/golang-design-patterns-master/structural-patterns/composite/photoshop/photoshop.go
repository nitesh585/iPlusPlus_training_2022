package photoshop

import "fmt"

// Point represents a point on the screen
type Point struct {
	// X is the x-coordinate
	X float64
	// Y is the y-coordinate
	Y float64
}

// Size represents a width and height size
type Size struct {
	// Width is the rectangle width
	Width float64
	// Height is the rectangle height
	Height float64
}

// Rect represents an rectangle
type Rect struct {
	// Location of the rectangle
	Location Point
	// Size of the rectanlge sides
	Size Size
}

// Drawer draws shapes
type Drawer struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (d *Drawer) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Drawing ellipse in rect %v", r)
	return nil
}

// DrawRect draws rectangle
func (d *Drawer) DrawRect(r Rect) error {
	fmt.Printf("Drawing rect %v", r)
	return nil
}

// VisualElement that is drawn on the screen
type VisualElement interface {
	// Draw draws the visual element
	Draw(drawer *Drawer) error
}

// Layer contains composition of visual elements
type Layer struct {
	// Elements of visual elements
	Elements []VisualElement
}

// Draw draws a layer
func (layer *Layer) Draw(drawer *Drawer) error {
	for _, element := range layer.Elements {
		if err := element.Draw(drawer); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

// Square represents a square
type Square struct {
	// Location of the square
	Location Point
	// Side size
	Side float64
}

// Draw draws a square
func (square *Square) Draw(drawer *Drawer) error {
	return drawer.DrawRect(Rect{
		Location: square.Location,
		Size: Size{
			Height: square.Side,
			Width:  square.Side,
		},
	})
}

// Circle represents a circle shape
type Circle struct {
	// Center of the circle
	Center Point
	// Radius of the circle
	Radius float64
}

// Draw draws a circle
func (circle *Circle) Draw(drawer *Drawer) error {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return drawer.DrawEllipseInRect(rect)
}
