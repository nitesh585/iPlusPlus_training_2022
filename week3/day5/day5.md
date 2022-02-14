## Continue Go lang

### My example of Interface

use case: Hospital Management Systems

```go
package main

import "fmt"

type Bill interface {
	getRoomBill() int
	// getDoctorFee() int
	// getOtherCharges() int
}

type Patient struct {
	name         string
	roomType     string
	numberOfDays int
}

func (p *Patient) getRoomBill() int {
	thresholRoomAmount := 0
	roomRentPerDay := 0
	switch p.roomType {
	case "luxury":
		thresholRoomAmount = 5000
		roomRentPerDay = 800
		break
	case "premium":
		thresholRoomAmount = 10000
		roomRentPerDay = 2000
		break
	default:
		thresholRoomAmount = 1000
		roomRentPerDay = 300
	}

	return thresholRoomAmount + roomRentPerDay*p.numberOfDays
}

func main() {
	p := Patient{name: "Amit", roomType: "noremal", numberOfDays: 4}
	fmt.Printf("Room Bill of %s: Rs %d", p.name, p.getRoomBill())
}
```

### Why interface in Go?

Interfaces are too big of a topic to give an all-depth answer here, but some things to make their use clear.

Interfaces are a tool. Whether you use them or not is up to you, but they can make code clearer, shorter, more readable, and they can provide a nice API between packages, or clients (users) and servers (providers).

Yes, you can create your own struct type, and you can "attach" methods to it, for example:

```go
type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

func main() {
    c := Cat{}
    fmt.Println("Cat says:", c.Say())
    d := Dog{}
    fmt.Println("Dog says:", d.Say())
}
```

We can already see some repetition in the code above: when making both Cat and Dog say something. Can we handle both as the same kind of entity, as animal? Not really. Sure we could handle both as interface{}, but if we do so, we can't call their Say() method because a value of type interface{} does not define any methods.

There is some similarity in both of the above types: both have a method Say() with the same signature (parameters and result types). We can capture this with an interface:

```go
type Sayer interface {
    Say() string
}
```

The interface contains only the signatures of the methods, but not their implementation.

Note that in Go a type implicitly implements an interface if its method set is a superset of the interface. There is no declaration of the intent. What does this mean? Our previous Cat and Dog types already implement this Sayer interface even though this interface definition didn't even exist when we wrote them earlier, and we didn't touch them to mark them or something. They just do.

Interfaces specify behavior. A type that implements an interface means that type has all the methods the interface "prescribes".

Since both implement Sayer, we can handle both as a value of Sayer, they have this in common. See how we can handle both in unity:

```go
animals := []Sayer{c, d}
for _, a := range animals {
    fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
}
```

(That reflect part is only to get the type name, don't make much of it as of now.)

The important part is that we could handle both Cat and Dog as the same kind (an interface type), and work with them / use them. If you were quickly on to create additional types with a Say() method, they could line up beside Cat and Dog:

```go
type Horse struct{}

func (h Horse) Say() string { return "neigh" }

animals = append(animals, Horse{})
for _, a := range animals {
    fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
}
```

Let's say you want to write other code that works with these types. A helper function:

```go
func MakeCatTalk(c Cat) {
    fmt.Println("Cat says:", c.Say())
}
```

Yes, the above function works with Cat and with nothing else. If you'd want something similar, you'd have to write it for each type. Needless to say how bad this is.

Yes, you could write it to take an argument of interface{}, and use type assertion or type switches, which would reduce the number of helper functions, but still looks really ugly.

The solution? Yes, interfaces. Simply declare the function to take a value of an interface type which defines the behavior you want to do with it, and that's all:

```go
func MakeTalk(s Sayer) {
    fmt.Println(reflect.TypeOf(s).Name(), "says:", s.Say())
}
```

You can call this function with a value of Cat, Dog, Horse or any other type not known until now, that has a Say() method. Cool.

Thanks @icza and @colm.anseo to answer on stackOverflow!
[more reading...](https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang)

## Must read resources

- [Go: Interfaces - I](https://golangbot.com/interfaces-part-1/#:~:text=In%20Go%2C%20an%20interface%20is%20a%20set%20of,the%20type%20decides%20how%20to%20implement%20these%20methods.)
- [Best Practices for Interfaces in Go](https://qvault.io/golang/golang-interfaces/#:~:text=Interfaces%20in%20Go%20allow%20us,unreadable%20and%20often%20buggy%20code.)
- [How do interfaces work in Go](https://www.calhoun.io/how-do-interfaces-work-in-go/)
- [Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go)
- [Debugging in VS code](https://code.visualstudio.com/Docs/editor/debugging)
