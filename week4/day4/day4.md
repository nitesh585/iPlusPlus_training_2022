## Continue with GO...

### Go routine

![go routine](https://miro.medium.com/max/552/1*GWYUFH14uOVLNHY-L1tv2w.jpeg)
Concurrency is a program’s ability to run more than one task independently in overlapping periods. In a concurrent program, several tasks can run at the same time in no particular order, which communicate, share resources, and interfere with each other.
Golang provides goroutines to support concurrency in Go. A goroutine is a function that executes simultaneously with other goroutines in a program and are lightweight threads managed by Go.

A goroutine takes about 2kB of stack space to initialize. In contrast, a standard thread can take up to 1MB, meaning creating a thousand goroutines takes significantly fewer resources than a thousand threads.

### Creating goroutines in Golang

Adding the keyword `go` in front of a function call executes the Go runtime as a goroutine.

```go
//Sequential way....
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func randSleep(name string, limit int){
    for i := 1; i <= limit; i++ {
        fmt.Println(name, rand.Intn(i))
        time.Sleep(time.Second)
    }
}

func main() {
    randSleep("first:", 4)
    randSleep("second:", 4)
}

// OUTPUT
// first: 0
// first: 1
// first: 2
// first: 3
// second: 0
// second: 0
// second: 1
// second: 0

```

In this sequential run, Go prints the numbers in the order the function calls. In the following program, the functions run concurrently:

```go
//Using Goroutine way....
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func randSleep(name string, limit int){
    for i := 1; i <= limit; i++ {
        fmt.Println(name, rand.Intn(i))
        time.Sleep(time.Second)
    }
}

func main() {
    go randSleep("first:", 4)
    go randSleep("second:", 4)
}

// OUTPUT
```

This program will not print anything out in the terminal because the main function completes before the goroutines execute, which is an issue; you don’t want your main to complete and terminate before the goroutines complete their execution.

The program terminates after the function below the goroutine completes its execution, regardless of whether the goroutine completes or not.

To solve this issue, Golang provides `WaitGroup`.

### `WaitGroup` in Golang

`WaitGroup`, provided in the sync package, allows a program to wait for specified goroutines. These are sync mechanisms in Golang that block the execution of a program until goroutines in the `WaitGroup` completely execute, as shown below:

```go
//Using Goroutine with WaitGroup....
package main

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

// wg is the pointer to a waitgroup
func randSleep(wg *sync.WaitGroup, name string, limit int){

    for i := 1; i <= limit; i++ {
        fmt.Println(name, rand.Intn(i))
        time.Sleep(time.Second)
    }
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    go randSleep("first:", 4)
    go randSleep("second:", 4)
    wg.wait()
}

// OUTPUT
// second: 0
// first: 0
// first: 1
// second: 1
// second: 1
// first: 0
// first: 1
// first: 0
// first: 4
// first: 1
// first: 6
// first: 7
// first: 2
```

`wg := new(sync.WaitGroup)` creates a new `WaitGroup` while `wg.Add(2)` informs WaitGroup that it must wait for two goroutines.

This is followed by `defer wg.Done()` alerting the `WaitGroup` when a goroutine completes.

`wg.Wait()` then blocks the execution until the goroutines’ execution completes.

### Communicating between Goroutines

Go provides a way for bidirectional communication between two goroutines through channels.

You can create a channel by declaring or using the make function:

```go
package main

import (
    "fmt"
)

func main() {
    // creating a channel by declaring it
    var mychannel1 chan int
    fmt.Println(mychannel1)

    // creating a channel using make()

    mychannel2 := make(chan int)
    fmt.Println(mychannel2)

}
```

Bidirectional channels in Go are blocking, meaning that when sending data into a channel, Go waits until the data is read from the channel before execution continues:

```go
package main

import (
    "fmt"
    "sync"
)

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
    defer wg.Done()
    for i := 1; i <= stop; i++ {
        limitchannel <- i
    }

}

func readChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
    defer wg.Done()
    for i := 1; i <= stop; i++ {
        fmt.Println(<-limitchannel)
    }
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)
    limitchannel := make(chan int)

    defer close(limitchannel)

    go writeChannel(wg, limitchannel, 3)
    go readChannel(wg, limitchannel, 3)
    wg.Wait()
}
// OUTPUT

// 1
// 2
// 3
```

With `limitchannel <- i`, the value of `i` enters the channel. `fmt.Println(<-limitchannel)` then receives the channel’s value and prints it out.

However, note that the number of sending operations must be equal to the number of receiving operations because if you send data to a channel and don’t receive it elsewhere, you get a `fatal error: all goroutines are asleep - deadlock!`.

### Buffered channels

If you were wondering why you must always receive from a channel after sending, this is because Go does not have anywhere to store the values passed into the channel.

However, you can create a channel that stores several values, meaning sending data into that channel won’t block until you exceed the capacity:

```go
package main

import (
    "fmt"
    "sync"
)

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
    defer wg.Done()
    for i := 1; i <= stop; i++ {
        limitchannel <- i
    }

}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(1)
    limitchannel := make(chan int, 2)

    defer close(limitchannel)

    go writeChannel(wg, limitchannel, 2)
    wg.Wait()

    fmt.Println(<-limitchannel)
    fmt.Println(<-limitchannel)
}

// OUTPUT

// 1
// 2
```

### Unit testing in GoLang

we will use the testing package to do some unit testing in Go.
Characteristics of a Golang test function:

- The first and only parameter needs to be `t *testing.T`
- It begins with the word Test followed by a word or phrase starting with a capital letter.(usually the method under test i.e. `TestValidateClient`)
- Calls `t.Error` or `t.Fail` to indicate a failure (I called t.Errorf to provide more details)
- `t.Log` can be used to provide non-failing debug information
- Must be saved in a file named something_test.go such as: `addition_test.go`

```go
// sum.go
package main

func Sum(x int, y int) int {
    return x + y
}

func main() {
    Sum(5, 5)
}
```

```go
// sum_test.go
package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
```

### Test Table

The concept of "test tables" is a set (slice array) of test input and output values. Here is an example for the Sum function:

```go
package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
```

<b>Run test</b>

```bash
go test
```

<b>Statement coverage</b>

```bash
go test -cover
```

<b>Generating an HTML coverage report</b>

```bash
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```

### Benchmark functions in GoLang

Go offers benchmarking via the testing package. Benchmark tests the time of execution for an operation. For example, a complex function execution time or a simple function executed a million times can be considered for benchmarking.

<b>Example:</b>

```go
package main

import (
     "fmt"
     "testing"
)

func BenchmarkF(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("Hi")
    }
}
```

The `b.N` is a variable to be adjusted by Go.

Benchmarks produce realistic expectations for the runtime of a program. It sheds light on the amount of time it can take when doing some expensive operations. It also shows how many of an operation can be done in a specific timeframe. It is really a very important metric through which many things regarding performance can be obtained that are otherwise hard to get.

## Must read resources

- [Concurrency With Golang Goroutines](https://tutorialedge.net/golang/concurrency-with-golang-goroutines/)

- [Implement Concurrency With Goroutine](https://betterprogramming.pub/golang-how-to-implement-concurrency-with-goroutines-channels-2b78b8077984)

- [Waitgroups in GoLang](https://golangdocs.com/waitgroups-in-golang)

- [Go: testing package](https://pkg.go.dev/testing)
- [Unit testing in GoLang](https://golangdocs.com/unit-testing-in-golang#:~:text=Unit%20testing%20in%20GoLang%20Unit%20testing%20is%20essential,package%20to%20do%20some%20unit%20testing%20in%20Go.)

- [Benchmark functions in GoLang](https://golangdocs.com/benchmark-functions-in-golang)
