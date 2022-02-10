## Continue with Go

### Functions

A function is a group of statements that together perform a task. Every Go program has at least one function, which is main().
A function declaration tells the compiler about a function name, return type, and parameters.

```go
package main

import "fmt"

func main () {

    fmt.Println("5 + 4 = ", add(5,4))
    fmt.Println(subtract(1,2,3,4,5))

}

func add(a,b int) int {
    return a+b
}

func subtract(args ... int) {
    sub := 0

    for _, value := range args {
        sub -= value
    }

    return sub
}
```

### Array

An array is a data structure in programming that is used to containerize data of the same type.

```go
package main

import "fmt"

func main() {

    // An Array holds a fixed number of values of the same type

    var favNums2[5] float64

    favNums2[0] = 163
    favNums2[1] = 78557
    favNums2[2] = 691
    favNums2[3] = 3.141
    favNums2[4] = 1.618

    // You access the value by supplying the index number

    fmt.Println(favNums2[3])

    // Another way of initializing an array

    favNums3 := [5]float64 { 1, 2, 3, 4, 5 }

    // How to iterate through an array (Use _ if a value isn't used)

    for i, value := range favNums3 {

    fmt.Println(value, i)

    }

    // Slices are like arrays but you leave out the size

    numSlice := []int {5,4,3,2,1}

    // You can create a slice by defining the first index value to
    // take through the last

    numSlice2 := numSlice[3:5] // numSlice3 == [2,1]

    fmt.Println("numSlice2[0] =", numSlice2[0])

    // If you don't supply the first index it defaults to 0
    // If you don't supply the last index it defaults to max

    fmt.Println("numSlice[:2] =", numSlice[:2])

    fmt.Println("numSlice[2:] =", numSlice[2:])

    // You can also create an empty slice and define the data type,
    // length (receive value of 0), capacity (max size)

    numSlice3 := make([]int, 5, 10)

    // You can copy a slice to another

    copy(numSlice3, numSlice)

    fmt.Println(numSlice3[0])

    // Append values to the end of a slice

    numSlice3 = append(numSlice3, 0, -1)

    fmt.Println(numSlice3[6])

}
```

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

### For loop

The for loop in Go works just like other languages. The loop starts with the keyword for. Then it initializes the looping variable then checks for condition, and then does the postcondition.

<b>Declaring a for-loop</b>

```go
package main

import "fmt"

func main() {
    for i:=0; i<3; i++ {  // start of the execution block
        fmt.Println("Hello")
    }
}
```

<b>infinite for-loop</b>

```go
for {
        // this block executes infinitely
}
```

<b>conditional for-loop</b>
Go don't have while loops. This variation only checks the condition and runs if it matches.

```go
package main

import "fmt"

func main() {
    i := 0
    for i<5 {
        fmt.Println(i)       // prints 0 2 4
        i += 2
    }
}
```

<b> range-for loop</b>
It makes the code much simpler.

```go
package main

import "fmt"

func main() {

    var items []int = []int{1, 2, 3, 4, 5}
    for i, v := range items {
        fmt.Println(i, v)
    }
}
```

### Map

Map which maps unique keys to values. A key is an object that you use to retrieve a value at a later date. Given a key and a value, you can store the value in a Map object.

```go
package main

import "fmt"

func main() {
    // A Map is a collection of key value pairs
    // Created with varName := make(map[keyType] valueType)

    presAge := make(map[string] int)

    presAge["Ram"] = 42

    fmt.Println(presAge["Ram"])

    // Get the number of items in the Map

    fmt.Println(len(presAge))

    // The size changes when a new item is added

    presAge["Rahul"] = 43
    fmt.Println(len(presAge))

    // We can delete by passing the key to delete

    delete(presAge, "Rahul")
    fmt.Println(len(presAge))

}
```

### Structure

Go allow you to define variables that can hold several data items of the same kind. A structure is another user-defined data type available in Go programming, which allows you to combine data items of different kinds. Structures are used to represent a record.

```go
package main

import "fmt"

// STRUCTS
type Rectangle struct{
    height float64
    width float64
}

func (rect *Rectangle) area() float64{
    return rect.width * rect.height
}

func main() {

    rect1 := Rectangle{height: 10, width: 10}

    fmt.Println("Rectangle is", rect1.width, "wide")

    fmt.Println("Area of the rectangle =", rect1.area())

}
```

## Must read resources

- [Go tutorial-1](https://github.com/callicoder/golang-tutorials)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)
- [Go maps in action](https://go.dev/blog/maps)
- [Go - Functions in Go language](https://ravichaganti.com/blog/get-set-go-functions-in-go-language/#named-return-values)
- [Exercism Go](https://exercism.org/tracks/go)
- [A Tour of Go](https://go.dev/tour/welcome/1)
