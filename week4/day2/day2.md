## Continue with Go

### Default Zero Value of all Types in Go

<b>Integer</b>

```go
package main

import "fmt"

func main() {
    var a int
    fmt.Print("Default zero value of int: ")
    fmt.Println(a)

    var b uint
    fmt.Print("Default zero value of uint: ")
    fmt.Println(b)
}
```

<b>Output:</b>

Default zero value of int: 0
Default zero value of uint: 0

<b>Float</b>

```go
package main

import "fmt"

func main() {
    var a float32
    fmt.Print("Default zero value of float: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of float: 0

<b>Complex Number</b>

```go
package main

import "fmt"

func main() {
    var a complex64
    fmt.Print("Default zero value of complex: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of complex: (0+0i)

<b>Byte</b>

```go
package main

import "fmt"

func main() {
    var a byte
    fmt.Print("Default zero value of byte: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of byte: 0

<b>Rune</b>

```go
package main

import "fmt"

func main() {
    var a rune
    fmt.Print("Default zero value of rune: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of rune: 0

<b>String</b>

```go
package main

import "fmt"

func main() {
    var a string
    fmt.Print("Default zero value of string: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of string: //An emtpy string. Hence nothing is printed

<b>Bool</b>

```go
package main

import "fmt"

func main() {
    var a bool
    fmt.Print("Default zero value of bool: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of bool: false

<b>Array</b>
Default value of an array is the default value of its values. For eg in the below code, there is an array of two lengths of type bool. When we print the output is [false false]

```go
package main

import "fmt"

func main() {
    var a [2]bool
    fmt.Print("Default Zero Value of an array: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default Zero Value of a array: [false false]

<b>Struct</b>
Default value of a struct is the default value of its field. For eg in the below code, there is a struct sample with two field. One is of int type and the other is of the bool type. We create an instance of this struct and when we print it, the output is {0 false}

```go
package main

import "fmt"

func main() {
    type sample struct {
        a int
        b bool
    }

    a := sample{}
    fmt.Print("Default Zero Value of an struct: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default Zero Value of a struct: {0 false}

<b>Map</b>
Default value of a map is nil. That is why the output of fmt.Println(a==nil) is true. When a map is passed to fmt.Println , it tries to print values in the map. That is why the output is map[]

```go
package main

import "fmt"

func main() {
    var a map[bool]bool
    fmt.Println(a == nil)
    fmt.Print("Printing map: ")
    fmt.Println(a)
}
```

<b>Output:</b>
true
Printing map: map[]

<b>Channel</b>

```go
package main

import "fmt"

func main() {
    var a chan int
    fmt.Print("Default zero value of channel: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of channel: <nil>

<b>Interface</b>

```go
package main

import "fmt"

func main() {
    var a chan interface{}
    fmt.Print("Default zero value of interface: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default zero value of interface: < nil>

<b>Slice</b>
Default value of a slice is nil. That is why the output of fmt.Println(a==nil) is true. When a slice is passed to fmt.Println , it tries to print values in the slice. That is why the output is []

```go
package main

import "fmt"

func main() {
    var a []int
    fmt.Println(a == nil)
    fmt.Print("Printing slice: ")
    fmt.Println(a)
}
```

<b>Output:</b>
true
Printing slice: []

<b>Function</b>

```go
package main

import "fmt"

func main() {
    var a func()
    fmt.Print("Default Zero Value of a func: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default Zero Value of a func: <nil>

<b>Pointer</b>

```go
package main

import "fmt"

func main() {
    var a \*int
    fmt.Print("Default Zero Value of a pointer: ")
    fmt.Println(a)
}
```

<b>Output:</b>
Default Zero Value of a pointer: <nil>
