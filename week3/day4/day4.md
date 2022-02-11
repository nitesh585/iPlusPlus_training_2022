## Continue with Go

### Type Casting

Type-casting means converting one type to another.Any type can be converted to another type but that does not guarantees that the value will remain intact or in fact preserved at all.

Unlike other languages, Go doesn’t support implicit type conversion. Although when dividing numbers, implicit conversions happen depending on the scenario. So we need to be very careful about what type to use where.

The syntax for general type casting is pretty simple. just use that other type name as a function to convert that value.

```go
package main

import (
    "fmt"
)

func main() {
    var a int = 42
    f := float64(a)      // convert int to float
    fmt.Println(f)       // 42
}
```

<b>Drawbacks of type-casting in GoLang</b>
Type conversion is a great way to change the original data type. But unfortunately, in some cases, it loses precision. It should be done sparingly. The needs of type conversion lie in the fact Go doesn’t support implicit type conversions. So in many cases, it should be done separately.

### While loop

in Go, there is no loop called while. There are only for-loops. The while loops can be emulated using the for-loops in Go.

```go
i := 0
for i < 10 {                      // emulates while (i < 10) {}
        // do something
}
```

### Do-while loop

The do-while loops can be emulated as well using just the for-loop. Here is an example showing just that.

```go
for {
    // do something
    if !condition {      // the condition stops matching
            break        // break out of the loop
    }
}

```

### My Example-1: struct with methods

```go
package main

import "fmt"

type Address struct {
	houseNo  int
	streetNo int
	city     string
	state    string
	country  string
}

type Passenger struct {
	id        int
	firstName string
	lastName  string
	address   Address
}

func (passenger Passenger) to_string() {
	fmt.Printf("Name: %s %s\n", passenger.firstName, passenger.lastName)
	fmt.Printf("Address: %d %d, %s, %s, %s \n", passenger.address.houseNo, passenger.address.streetNo, passenger.address.city, passenger.address.state, passenger.address.country)
}

func (passenger Passenger) getFullName() string {
	return passenger.firstName + " " + passenger.lastName
}

func main() {
	address := Address{55, 6, "New Delhi", "Delhi", "India"}
	passenger := Passenger{2011, "Nitesh", "Yadav", address}
	passenger.to_string()
	fmt.Println(passenger.getFullName())
}
```

### My Example-2: struct with init(constructor) and methods

```go
package main

import "fmt"

type Doctor struct {
	id              int
	name            string
	patientAssigned int
	capacity        int
	available       bool
}

func (e *Doctor) init(id, patientAssigned, capacity int, name string) {
	e.id = id
	e.name = name
	e.capacity = capacity
	e.patientAssigned = patientAssigned
	if capacity <= patientAssigned {
		e.available = false
	} else {
		e.available = true
	}
}

func (e *Doctor) isAvailable() {
	if e.available {
		fmt.Printf("%s is available\n", e.name)
	} else {
		fmt.Printf("%s is not available\n", e.name)
	}
}

func main() {
	////--------------------Hospital management
	doctor1 := new(Doctor)
	doctor2 := new(Doctor)

	doctor1.init(2011, 3, 5, "Amit")
	doctor2.init(2012, 6, 5, "Rohit")

	doctor1.isAvailable()
	doctor2.isAvailable()
}

```

### interface

An interface is an abstract concept which enables polymorphism in Go. A variable of that interface can hold the value that implements the type. Type assertion is used to get the underlying concrete value
An interface is declared as a type. Here is the declaration that is used to declare an interface.

`type interfaceName interface{}`

```go

package main

import (
    "fmt"
)

type Person interface{
    greet() string
}

type Human struct{
    Name string
}

func (h *Human)greet() string {
    return "Hi, I am " + h.Name
}

func isAPerson(h Person) {
    fmt.Println(h.greet())
}

func main() {
    var a = Human{"John"}

    fmt.Println(a.greet())    // Hi, I am John

    // below function will only work
    // if a is also a person.
    // Here we can see polymorphism in action.
    isAPerson(&a)             // Hi, I am John
}

```

## Must read resources

- [Go: Struct](https://golangbot.com/structs/)
- [Construct For Loops](https://www.digitalocean.com/community/tutorials/how-to-construct-for-loops-in-go)
- [Interfaces in GoLang](https://golangdocs.com/interfaces-in-golang)
- [Go tutorial-1](https://github.com/callicoder/golang-tutorials)
- [Exercism Go](https://exercism.org/tracks/go)
- [A Tour of Go](https://go.dev/tour/welcome/1)
