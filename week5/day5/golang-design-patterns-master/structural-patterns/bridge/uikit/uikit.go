package uikit

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

// Drawer draws on the underlying graphics device
type Drawer interface {
	// DrawEllipseInRect draws an ellipse in rectanlge
	DrawEllipseInRect(Rect) error
}

// OpenGL drawer
type OpenGL struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (gl *OpenGL) DrawEllipseInRect(r Rect) error {
	fmt.Printf("OpenGL is drawing ellipse in rect %v", r)
	return nil
}

// Direct2D drawer
type Direct2D struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (d2d *Direct2D) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Direct2D is drawing ellipse in rect %v", r)
	return nil
}

// Circle represents a circle shape
type Circle struct {
	// DrawingContext for this circle
	DrawingContext Drawer
	// Center of the circle
	Center Point
	// Radius of the circle
	Radius float64
}

// Draw draws a circle
func (circle *Circle) Draw() error {
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

	return circle.DrawingContext.DrawEllipseInRect(rect)
}
