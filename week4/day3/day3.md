## Continue with Go

### Pointer and Reference

<b>Pointers:</b> A pointer is a variable that holds memory address of another variable. A pointer needs to be dereferenced with \* operator to access the memory location it points to.

<b>References:</b> A reference variable is an alias, that is, another name for an already existing variable. A reference, like a pointer, is also implemented by storing the address of an object.

```go
package main

import (
 "fmt"
)

func main() {
    item := ""
    passByValue(item)
    fmt.Println(item)
    passByReference(&item)
    fmt.Println(item)
}

func passByValue(item string) {
    item = "hello"
}

func passByReference(item *string) {
    *item = "world"
}
```

### Scope of variables

A scope in Golang is the region of a program where a defined variable can exist, and beyond that, a variable cannot be accessed.
file handling

There are three places where the variables can be declared in Golang.

- Local variables
- Global variables
- Formal parameters

<b>Local variables in Golang</b>
Local variables are those declared inside a function or a block. Local variables can be used only by statements that are inside that function or block of code.

<b>Global variables in Golang</b>
The variables which are defined outside of the function or a block is termed as Global variables. These are available throughout the lifetime of a program.

<b>Formal Parameters in Golang</b>
Formal parameters are treated as the local variables within that function, and they take preference over the global variables.

```go
// hello.go

package main

import "fmt"

/* global variable declaration */
var x int = 20

func main() {
    /* local variable declaration in main function */
    var x int = 10
    var y int = 20
    var z int

    fmt.Printf("value of x in main() function = %d\n", x)
    z = sum(x, y)
    fmt.Printf("value of z in main() function = %d\n", z)
}

/* function to add two integers */
func sum(a, b int) int {
    fmt.Printf("value of a in sum() function = %d\n", a)
    fmt.Printf("value of b in sum() function = %d\n", b)

    return a + b
}
```

### File Handling

<b>Reading files in GoLang</b>
Reading files is one of the most essential tasks in programming. It allows us to see content, modify and write it using programming. In this post, we will see how to read file contents in Go.

```go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    content, err := ioutil.ReadFile("text.txt")     // the file is inside the local directory
    if err != nil {
        fmt.Println("Err")
    }
    fmt.Println(string(content))    // This is some content
}
```

<b>Write files in GoLang</b>
Go has the ioutil package which makes writing a file a lot easier. So, let’s see what are the
ways we can do it.

```go
package main

import (
    "fmt"
    "os"
    //"io/ioutil"
)

func main() {

  // create the file
  f, err := os.Create("test.txt")
  if err != nil {
    fmt.Println(err)
  }
  // close the file with defer
  defer f.Close()

  // do operations

  //write directly into file
  f.Write([]byte("a string"))

  // write a string
  f.WriteString("\nThis is a pretty long string")

  // // write from a specific offset
  f.WriteAt([]byte("another string"), 42)     // 12 is the offset from start
}
```

### My Example

Use case: Train Ticket Booking System

```json
//details.json
[
  { "id": 121, "name": "ABA" },
  { "id": 122, "name": "ABB" },
  { "id": 123, "name": "ABC" },
  { "id": 124, "name": "ABD" },
  { "id": 125, "name": "ABE" }
]
```

```go
//main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Train struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t Train) getName() string {
	return t.Name
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		panic("Error when opening file: ")
	}
}

func getTrainDetails(id int) (Train, bool) {
	content, err := ioutil.ReadFile("details.json")
	check(err)

	// Now let's unmarshall the data into `payload`
	var data []Train
	err = json.Unmarshal(content, &data)
	check(err)

	for _, train := range data {
		if train.ID == id {
			return train, true
		}
	}

	return Train{}, false
}

func main() {
	// fmt.Print("hello")
	train, ok := getTrainDetails(125)
	if ok {
		fmt.Printf("Train ID found!\nTrain name: %s", train.getName())
	} else {
		fmt.Println("Train ID not found!")
	}
}
```

## Must read resouces

- [Go: Global variables](https://golangdocs.com/golang-global-variables)

- [Pointers and Passby value/reference in GOlang](https://medium.com/wesionary-team/pointers-and-passby-value-reference-in-golang-a00c8c59b7f1)

- [Reading files in GoLang](https://golangdocs.com/reading-files-in-golang)
- [Reading files in GoLang](https://golangdocs.com/write-files-in-golang)
- [Delete files in GoLang](https://golangdocs.com/delete-files-in-golang)

- [Optional: Read files](https://golangbot.com/read-files/)
- [Optional: Read and write file in Golang / File handling in Golang](https://learnetutorials.com/golang/file-handling#:~:text=%20Read%20and%20write%20file%20in%20Golang%20%2F,call%20the%20os.Remove%20%28%29%20function%2C%20enter...%20More%20)
