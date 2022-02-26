#### Introduction

The Bridge Pattern is a creational design pattern used to decouple an
abstraction from its implementation. The bridge uses encapsulation,
aggregation, and can use inheritance to separate responsibilities into
different classes.

#### Purpose

- Decouple an abstraction from its implementation that allows both to vary independently.
- Publish interface in an inheritance hierarchy, and bury implementation in its own inheritance hierarchy.

#### Design Pattern Diagram

The objects participating in this pattern are presented on the following diagram:

![alt tag](http://blog.ralch.com/media/golang/design-patterns/bridge.gif)

- `Abstraction` defines the abstraction's interface and maintains a reference to an object of type `Implementor`
- `RefinedAbstraction` extends the interface defined by `Abstraction`
- `Implementor` defines the interface for implementation objects
- `ConcreteImplementor` implements the `Implementor` interface and defines its concrete implementation

#### Implementation

Consider building an UI package that supports drawing different shapes on the
screen by supporting `OpenGL` and `Direct2D` rendering technologies. In our
particular example should be able to draw a circle by supporting both rendering
systems.

In order to do that we should separate the `Circle` struct from its drawing
implementation:

```Golang
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
```

The `Drawer` defines the contract between the `Circle` abstraction and its
implementation:

```Golang
// Drawer draws on the underlying graphics device
type Drawer interface {
	// DrawEllipseInRect draws an ellipse in rectanlge
	DrawEllipseInRect(Rect) error
}
```

For every of the supported rendering technologies we should implement a struct
that obeys the `Drawer` interface:

This is a sample implementation of `OpenGL` drawer:

```Golang
// OpenGL drawer
type OpenGL struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (gl *OpenGL) DrawEllipseInRect(r Rect) error {
	fmt.Printf("OpenGL is drawing ellipse in rect %v", r)
	return nil
}
```

This is a sample implementation of `Direct2D` drawer:

```Golang
// Direct2D drawer
type Direct2D struct{}

// DrawEllipseInRect draws an ellipse in rectangle
func (d2d *Direct2D) DrawEllipseInRect(r Rect) error {
	fmt.Printf("Direct2D is drawing ellipse in rect %v", r)
	return nil
}
```

Then we can easily render a circle by using the desired drawing system. We should
just change the implementation by setting `DrawingContext` field:

```Golang
openGL := &uikit.OpenGL{}
direct2D := &uikit.Direct2D{}

circle := &uikit.Circle{
	Center: uikit.Point{X: 100, Y: 100},
	Radius: 50,
}

circle.DrawingContext = openGL
circle.Draw()

circle.DrawingContext = direct2D
circle.Draw()
```

#### Verdict

The Bridge Pattern is designed up-front to let the abstraction and the
implementation vary independently. That allows us to independently change the
implementation from its abstraction.
