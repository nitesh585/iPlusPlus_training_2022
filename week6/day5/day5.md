## Gin Gonic Framework

![gin-gonic](https://th.bing.com/th/id/OIP.RSU5mKZnf5L2fwmJ3xRdYAAAAA?pid=ImgDet&rs=1)
Gin is a high-performance HTTP web framework written in Golang (Go). Gin allows you to build web applications and microservices in Go. It contains a set of commonly used functionalities (eg, routing, middleware support, rendering, etc.) that reduce boilerplate code and make it easier to build web applications.

### Features

- <b>Fast</b>
  Radix tree based routing, small memory foot print. No reflection. Predictable API performance.

- <b>Middleware support</b>
  An incoming HTTP request can be handled by a chain of middlewares and the final action. For example: Logger, Authorization, GZIP and finally post a message in the DB.

- <b>Crash-free</b>
  Gin can catch a panic occurred during a HTTP request and recover it. This way, your server will be always available. As an example - it’s also possible to report this panic to Sentry!

- <b>JSON validation</b>
  Gin can parse and validate the JSON of a request - for example,checking the existence of required values.

- <b>Routes grouping</b>
  Organize your routes better. Authorization required vs non required, different API versions… In addition, the groups can be nested unlimitedly without degrading performance.

- <b>Error management</b>
  Gin provides a convenient way to collect all the errors occurred during a HTTP request. Eventually, a middleware can write them to a log file, to a database and send them through the network.

- <b>Rendering built-in</b>
  Gin provides an easy to use API for JSON, XML and HTML rendering.

- <b>Extendable</b>
  Creating a new middleware is so easy, just check out the sample codes.

### Installation

To install Gin package, you need to install Go and set your Go workspace first.

Download and install it:

```sh
$ go get -u github.com/gin-gonic/gin
```

Import it in your code:

```go
import "github.com/gin-gonic/gin"
```

(Optional) Import `net/http`. This is required for example if using constants such as `http.StatusOK`.

```go
import "net/http"
```

Create your project folder and cd inside

```sh
$ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$\_"
```

Copy a starting template inside your project

```sh
$ curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

Run your project

```sh
$ go run main.go
```

### Getting Started

First, create a file called example.go:

```sg
# assume the following codes in example.go file
$ touch example.go
```

Next, put the following code inside of `example.go`:

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c \*gin.Context) {
        c.JSON(200, gin.H{
        "message": "pong",
    })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```

And, You can run the code via `go run example.go`:

```sh
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```

## DynamoDB

![dyanamoDB](https://user-images.githubusercontent.com/6509926/70553550-f033b980-1b40-11ea-9192-759b3b1053b3.png)

Amazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable performance with seamless scalability. DynamoDB lets you offload the administrative burdens of operating and scaling a distributed database so that you don't have to worry about hardware provisioning, setup and configuration, replication, software patching, or cluster scaling. DynamoDB also offers encryption at rest, which eliminates the operational burden and complexity involved in protecting sensitive data.

## Must read resources:

- [Amazon DynamoDB: How It Works](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.html)
- [Setting Up DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SettingUp.html)
- [Accessing DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/AccessingDynamoDB.html)
- [Gin Web Framework](https://gin-gonic.com/)
- [Github - Gin-Gonic](https://github.com/gin-gonic/gin)
