## Monolithic Vs Micro Service Architecture

![comparison image](https://blog.kohactive.com/content/images/2018/11/monolith_microservices.jpg)

| Key                                 | Monolithic architecture                                                                                                             | Microservices architecture                                                                          |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| Basic                               | Monolithic architecture is built as one large system and is usually one code-base                                                   | Microservices architecture is built as small independent module based on business functionality     |
| Scale                               | It is not easy to scale based on demand                                                                                             | It is easy to scale based on demand.                                                                |
| Deployment                          | Large code base makes IDE slow and build time gets increase.                                                                        | Each project is independent and small in size. So overall build and development time gets decrease. |
| Tightly Coupled and Loosely coupled | It extremely difficult to change technology or language or framework because everything is tightly coupled and depend on each other | Easy to change technology or framework because every module and project is independent              |
| Database                            | It has shared database                                                                                                              | Each project and module has their own database                                                      |

## Intro to Go Lang

GoLang is the only language that incorporates all three sought-after capabilities. Namely, ease of coding, efficient code-compilation and efficient execution. To be precise, Go came to offer the following solutions:

- Fast-paving compilation and execution
- A boost to code readability and documentation
- Offering a thoroughly consistent language
- Facilitating easy versioning of the program

<b>Hello World: First Golang Programme</b>

```go
package main

import "fmt"

func main () {
    fmt.Println("Hello World!")
}
```

<b>Variables and Constants</b>
Variable is nothing but a name given to a storage area. Constants refer to fixed values that the program may not alter during its execution

```go
package main

import "fmt"

func main () {

    var x int = 5 //declaration of variable x

    var (
    a = 10
    b = 15 //declaration of multiple variables
    )

    y := 7 // shorthand declaration of variables

    const pi float64 = 3.14272345 //declaration of constants

    var name string = "Aryya Paul" //declaration of a string
}
```

<b>Decision Making</b>
Decision making is a critical part of programming. In Golang, we can implement decision making using the ‘if-else’ and ‘switch’ keywords

```go
package main

import "fmt"

func main() {

    // If Statement

    age := 15

    if age >= 16 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Not an adult")
    }

    // You can use else if perform different actions, but once a match
    // is reached the rest of the conditions aren't checked

    if age >= 16 {
        fmt.Println("in school")
    } else if age >= 18 {
        fmt.Println("in college")
    } else {
        fmt.Println("probably dead")
    }

    // Switch statements are used when you have limited options

    switch age {
        case 16: fmt.Println("Go Drive")
        case 18: fmt.Println("Go Vote")
        default: fmt.Println("Go Have Fun")
    }

}
```

## Must read resources

- [Monolithic Vs Microservices](https://articles.microservices.com/monolithic-vs-microservices-architecture-5c4848858f59)
- [Install Go](https://go.dev/doc/install)
- [Exercism Go](https://exercism.org/tracks/go)
- [A Tour of Go](https://go.dev/tour/welcome/1)
- [Get started with Go](https://go.dev/doc/tutorial/getting-started)
