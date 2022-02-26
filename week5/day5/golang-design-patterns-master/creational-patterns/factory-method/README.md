#### Introduction

The `Factory Method` pattern is a design pattern used to define a runtime
interface for creating an object. It’s called a factory because it creates
various types of objects without necessarily knowing what kind of object it
creates or how to create it.

#### Purpose

- Allows the sub-classes to choose the type of objects to create at runtime
- It provides a simple way of extending the family of objects with minor
	changes in application code.
- Promotes the loose-coupling by eliminating the need to bind
	application-specific structs into the code

#### Desing Pattern Diagram

The structs and objects participating in this pattern are: product,
concreate product, creator and concrete creator. The Creator contains one
method to produce one type of product related to its type.

![alt tag](http://blog.ralch.com/media/golang/design-patterns/factory-method.gif)

- `Product` defines the interface of objects the factory method creates
- `ConcreteProduct` implements the Product interface
- `Creator` declares the factory method, which returns an object of type Product
- `ConcreteCreator` overrides the factory method to return an instance of a Concrete Product

#### Implementation

The Factory Method defines an interface for creating objects, but lets
subclasses decide which classes to instantiate. In these example, we will adopt
the pattern to create document object model of Scalable Vector Graphics.

The SVG format can contains multiple elements. In this example, we will illustrate
only some of the shape elements. In the context of `Factory Method` design pattern,
they are our product.

Every shape implements the `Shape` interface, which expose a `Draw` function that
generates the required XML element:

```Golang
type Shape interface {
	Draw(io.Writer) error
}
```

In the following code snippets, we will illustrate two implementations of `Shape` 
interface `Circle` and `Ractangle`:

```Golang
type Circle struct {
	Location Point
	Radius   float64
}

func (c *Circle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<circle cx="%f" cy="%f" r="%f"/>`, c.Location.X, c.Location.Y, c.Radius)
	return err
}
```

```Golang
type Rectangle struct {
	Location Point
	Size     Size
}

func (rect *Rectangle) Draw(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<rect x="%f" y="%f" width="%f" height="%f"/>`, rect.Location.X, rect.Location.Y, rect.Size.Width, rect.Size.Height)
	return err
}
```

Every of them has a function that is responsible for their instantiation based
on the provided `Viewport`. The `Viewport` is an argument which keeps an information
about the location and the size of the view port.

```Golang
type ShapeFactory interface {
	Create(viewport Viewport) Shape
}
```

The `CircleFactory` creates a `Circle` instance that has radius, which fits
the viewport:

```Golang
type CircleFactory struct{}

func (factory *CircleFactory) Create(viewport Viewport) Shape {
	return &Circle{
		Location: viewport.Location,
		Radius:   math.Min(viewport.Size.Width, viewport.Size.Height),
	}
}
```

The `RectangleFactory` produces a rectangle that fits the viewport:

```Golang
type RactangleFactory struct{}

func (factory *RactangleFactory) Create(viewport Viewport) Shape {
	return &Rectangle{
		Location: viewport.Location,
		Size:     viewport.Size,
	}
}
```

The main object `Document` has a `Draw` function, which composes a different
shapes created by provided factories. The `Document` can be instaciated with
different set of factories. This allow to customize and change the document's
content:

```Golang
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
```

We should instaciate the `Document` struct with the available factories in
the following way:

```Golang
doc := &svg.Document{
	ShapeFactories: []svg.ShapeFactory{
		&svg.CircleFactory{},
		&svg.RactangleFactory{},
	},
}

doc.Draw(os.Stdout)
```

Important aspects when we implement the Factory Method design pattern are:

- Designing the arguments of the factory method
- Considering an internal object pool that will allow object cache and reuse instead
of created from scratch

#### Verdict

The Factory Method is one of the most used design patterns. It makes a design
more customizable and only a little more complicated. Other design patterns
require new structs, whereas Factory Method only requires a new operation.
The Factory Method is similar to Abstract Factory but without the emphasis on
families.

