#### Introduction

A Composite Design Pattern is a structural pattern that uses to employ
composition when implementing an interface rather than using multiple
inheritance. It composes objects into tree structures and lets
clients treat individual objects and compositions uniformly.

There are situations when clients ignore the difference between
compositions of objects and individual objects. If dealing with multiple
objects in the same way as handle each of deam is identical, the Composite
Design pattern is a good choice to decrease the complexity and treat them as
homogeneous.

#### Purpose

- The intent of this pattern is to compose objects into tree structures to
	represent part-whole hierarchies.
- Composite lets clients treat individual objects and compositions of objects
	uniformly.

#### Design Pattern Diagram

The Composite Pattern provides the following units: Component, Leaf and
Composite.

![alt tag](http://blog.ralch.com/media/golang/design-patterns/composite.gif)

- `Component` is an interface for all components, including composite ones. It declares the interface for objects in the composition
- `Leaf` represents leaf objects in the composition implements all `Component` methods
- `Composite` represents a composite `Component` that has children. Usually
	implements all Componenet methods and methods to manipulate children.

#### Implementation

The Composite Design Pattern is very common approach for implementing a
document object model hierarchy. Such an example are image editors, which
compose different shapes and layers into hierarchy.

Lets implement a basic architecture of such an editor. Lets use the following
interface:

```Golang
// VisualElement that is drawn on the screen
type VisualElement interface {
	// Draw draws the visual element
	Draw(drawer *Drawer) error
}
```

The editor supports two kind of shapes (circle and square). Each of the structs
that represents the coresponding shape obeys the `VisualElement` interface by
implementing a `Draw` function that has exactly the same signiture exposed in
the interface. The following code snippet illustrate the implementation of
thoes components:

```Golang
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
```

In order to allow composition and drawing of multiple shapes on the screen, a
`Layer` composite struct is implemented. It contains an array of
`VisualElement`. It is responsible to interate over the elements and draw each
of them. As you can see the actual struct uses the `VisualElement` interface as
a contract to support different shapes no matter what is their type.

```Golang
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
```

The object can be composed as it is shown in the following code snippet:

```Golang
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
```

#### Verdict

Dealing with Tree-structured data makes code more complex, and therefore,
error prone. The Composite Design Patterns provides a solution that allows
treating complex and primitive objects uniformly. The operations you can
perform on all the composite objects often have a least common relationship
that allows handling a set of object as a single unit.
