## RabbitMQ Vs Kafka

![kafka-vs-rabbitq](https://i.stack.imgur.com/KHpje.png)

## Design Patterns

In software engineering, a design pattern is a general repeatable solution to a commonly occurring problem in software design. A design pattern isn't a finished design that can be transformed directly into code. It is a description or template for how to solve a problem that can be used in many different situations.

## Microservice Architecture

![microservice-arc](https://res.cloudinary.com/practicaldev/image/fetch/s--seen3BGm--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://user-images.githubusercontent.com/2697570/49395813-cd094980-f737-11e8-9e9a-6c20db5720c4.jpg)

Microservices architecture (often shortened to microservices) refers to an architectural style for developing applications. Microservices allow a large application to be separated into smaller independent parts, with each part having its own realm of responsibility. To serve a single user request, a microservices-based application can call on many internal microservices to compose its response.

### Monolithic Vs Microservices

![mono-vs-micro](https://www.hitechnectar.com/wp-content/uploads/2020/02/Monolithic-architecture-vs-microservices-architecture-comparison.jpg)

## GO Examples:

### Worker pool

Worker pools is a design in which a fixed number of m workers (Go goroutines) works on n tasks in a work queue (Go channel). The work resides in a queue until a worker finish its current task and pull a new one.

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

### Ticker timer

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
```

### Atomic counter

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("ops:", ops)
}
```

## Must read resources

- [Pattern: Microservice Architecture](https://microservices.io/patterns/microservices.html)

- [Go by Example: Worker Pools](https://gobyexample.com/worker-pools)

- [Go by Example: Tickers](https://gobyexample.com/tickers)

- [Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)

- [Atomic Package](https://pkg.go.dev/sync/atomic)
