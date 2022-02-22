## Continue with Go...

### Select

![select-go-lang](https://559987-1802630-raikfcquaxqncofqfm.stackpathdns.com/assets/images/go/package/image8.png)

GoLang select statement is like the switch statement, which is used for multiple channels operation. This statement blocks until any of the cases provided are ready.
Select statements work like switch statements but instead of having concrete cases it has operations that are either sending or receiving using channels.

```go
package main

import (
    "fmt"
)

func g1(ch chan int) {
    ch <- 12
}

func g2(ch chan int) {
    ch <- 32
}

func main() {

    ch1 := make(chan int)
    ch2 := make(chan int)

    go g1(ch1)
    go g1(ch2)

    select {
    case v1 := <-ch1:
        fmt.Println("Got: ", v1)
    case v2 := <-ch2:
        fmt.Println("Got: ", v2)
    }
}
```

It is simply random. We cannot predict the output since select works in a very different way. It chooses any output if all of the statements are ready for execution.

## Mocking

Mocking in golang is done with the help of interfaces. Mocking in unit testing is important as it ensure that variables, methods and functions modified outside the scope of the function being tested do not affect the test output.

Followed the example of: Testing with GoMock: A Tutorial (in must read resources) and code present in `go-mocking` dir.

# Must read resources

- [Select statement in GoLang](https://golangdocs.com/select-statement-in-golang)
- [Select in Golang](https://learnetutorials.com/golang/select#:~:text=What%20is%20the%20difference%20between%20select%20and%20switch,Require%20a%20fallthrough%20%202%20more%20rows%20)

- [Testing with GoMock: A Tutorial](https://blog.codecentric.de/en/2017/08/gomock-tutorial/)

- [Golang Testing: Mocking and Error Handling](https://williaminfante.medium.com/golang-testing-mocking-and-error-handling-fbfe7f6008b9)
- [Mocking Methods in Go](https://dev.to/dmigwi/mocking-methods-in-go-5fg#:~:text=Mocking%20in%20golang%20is%20done%20with%20the%20help,test%20output.%20Here%20is%20the%20implementation%20of%20mocking.go)
