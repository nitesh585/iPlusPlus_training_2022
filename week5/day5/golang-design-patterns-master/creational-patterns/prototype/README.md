# The Prototype Design Pattern

#### Preface

`The Prototype Pattern` creates duplicate objects while keeping performance
in mind. It's a part of the creational patterns and provides one of the best
ways to create an object.

In the mainstream languages (like C# and JAVA), it requires implementing a
prototype interface which tells to create a clone of the current object. It is
used when creation of object directly is costly.

For instance, an object is to be created after a costly database operation. We
can cache the object, returns its clone on next request and update the database
as and when needed thus reducing database calls.

#### Purpose

- Specify the kind of objects to create using a prototypical instance, and
	create new objects by copying this prototype.

#### Desing Pattern Diagram

{{< figure src="/media/golang/design-patterns/prototype.gif" alt="Prototype Class Diagram" >}}

- `Prototype` declares an interface for cloning itself
- `ConcretePrototype` implements an operation for cloning itself
- `Client` creates a new object by asking a prototype to clone itself

#### Implementation

In Golang, the pattern is applicable only in situation that we want to
customize how the object is cloned. We will explore two examples regarding
both situations.

Lets build a system that generates a different configuration files depending on
our needs. In first place, we have a struct `Config` that looks like:

```Golang
package configurer

// Config provides a configuration of microservice
type Config struct {
	workDir string
	user    string
}

// NewConfig create a new config
func NewConfig(user string, workDir string) Config {
	return Config{
		user:    user,
		workDir: workDir,
	}
}

// WithWorkDir creates a copy of Config with the provided working directory
func (c Config) WithWorkDir(dir string) Config {
	c.workDir = dir
	return c
}

// WithUser creates a copy of Config with the provided user
func (c Config) WithUser(user string) Config {
	c.user = user
	return c
}
```

We want to be able to mutate the object without affecting its initial instance.
The goal is to be able to generate different configuration files without loosing
the flexibility of customizing them without mutation of the initial default
configuration.

As you can see the functions `WithWorkDir`, `WithUserID` and `WithGroupID` are
declared for the struct `Config` (not for `*Config`). At the time, when they are
called the object is copied by the `Golang` runtime. This allows us to modify it
without affecting the original object.

Lets see it's usage in action:

```Golang
config := configurer.NewConfig(10, 10, "/home/guest")
rootConfig := config.WithUserID(0).WithGroupID(0).WithWorkDir("/root")

fmt.Println("Guest Config", config)
fmt.Println("Root Config", rootConfig)
```

Now lets explore the classic implementation of `The Prototype Design Pattern`.
We will assume that we are developing again document object model for a custom
document format. The core object is an `Element` structure which has parent and
children.

```Golang
// Element represents an element in document object model
type Element struct {
	text     string
	parent   Node
	children []Node
}

// NewElement makes a new element
func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

// Parent returns the element parent
func (e *Element) Parent() Node {
	return e.parent
}

// SetParent sets the element parent
func (e *Element) SetParent(node Node) {
	e.parent = node
}

// Children returns the element children elements
func (e *Element) Children() []Node {
	return e.children
}

// AddChild adds a child element
func (e *Element) AddChild(child Node) {
	copy := child.Clone()
	copy.SetParent(e)
	e.children = append(e.children, copy)
}

// Clone makes a copy of particular element. Note that the element becomes a
// root of new orphan tree
func (e *Element) Clone() Node {
	copy := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}
	for _, child := range e.children {
		copy.AddChild(child)
	}
	return copy
}

// String returns string representation of element
func (e *Element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}
```

The contract that exposes the `Clone` funcion is the `Node` interface:

```Golang
// Node a document object model node
type Node interface {
	// Strings returns nodes text representation
	String() string
	// Parent returns the node parent
	Parent() Node
	// SetParent sets the node parent
	SetParent(node Node)
	// Children returns the node children nodes
	Children() []Node
	// AddChild adds a child node
	AddChild(child Node)
	// Clone clones a node
	Clone() Node
}
```

We want to extract a particular subtree of concrete element hierary. We want to
use it as independent document object model. In order to do that, we should use
the clone function:

```Golang
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
```

The sample above creates a tree from the subtree pointed by `engManagerNode`
variable.

You can get the full source code from
[github](https://github.com/svett/golang-design-patterns/tree/master/creational-patterns/prototype).

#### Verdict

One of the disadvantages of this pattern is that the process of copying an object
can be complicated. In addition, structs that have circular references to other
classes are difficult to clone. Its overuse could affect performance, as the
prototype object itself would need to be instantiated if you use a registry of
prototypes.

In the context of `Golang`, I don't see any particular reason to adopt it.
`Golang` already provides builtin mechanism for cloning objects.

