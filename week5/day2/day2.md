## Continue with Go....

### Go routine example: 

```go

package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 50; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers(pings)
	go alphabets(pings, pongs)
	time.Sleep(9000 * time.Millisecond)
	fmt.Println("main done")
}

```

### Simple channel example:

```go

package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

    ping(pings, "Hello")
    pong(pings, pongs)

    fmt.Print(<-pongs)

}

```

### Non-Blocking Channle Operations:

```go

package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("receive msg ", msg)
	default:
		fmt.Println("no receive msg ")
	}

	msg := "hi"
	// message change don't have buffe that's why default statement got printed
	select {
	case message <- msg:
		fmt.Println("receive msg ", msg)
	default:
		fmt.Println("no receive msg ")
	}

	select {
	case msg := <-message:
		fmt.Println("receive msg ", msg)
	case msg := <-signal:
		fmt.Println("receive sig ", msg)
	default:
		fmt.Println("no activity ")
	}

}

```
