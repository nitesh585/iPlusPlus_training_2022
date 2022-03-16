## HTTP Server Implementation Using Go

`net` package contains http package that provides both HTTP client (to make http requests) and HTTP server (listens to http requests) implementations.

The crux to understanding the http server implementation is understanding below things

- <b>Request</b> – it defines the request parameters i.e, Method, Api Signature, request headers, body, query params etc
- <b>Response</b> – defines the response parameters i.e, Status code, response body, headers
- <b>Pair of API signature and its handler</b> – Each API signature corresponds to a handler. You can think of handler as a function which is invoked when a request is made for that particular API signature. The mux registers these pairs of API signature and its handler
- <b>Mux</b>– It acts as a router. Depending upon API signature of the request, it routes the request to the registered handler for that API signature. The handler will handle that incoming request and provide the response . For eg an API call with “/v2/teachers” might be handled by a different function and API call with “/v2/students” might be handled by some other function. So basically based upon API signature( and also request method sometimes) , it decides which handler to invoke.
- <b>Listener</b> – It runs on the machine, which listens to a particular port. Whenever it receives the request on that port it forwards the request to the mux. It also handles other functionalities as well but we will not discuss those it in this article.

![mux](https://i1.wp.com/golangbyexample.com/wp-content/uploads/2020/07/http.jpg?w=712&ssl=1)

```go
//main.go

package main

import (
	"net/http"
)

func main() {

	//Create the default mux
	mux := http.NewServeMux()

	//Handling the /v1/teachers. The handler is a function here
	mux.HandleFunc("/v1/teachers", teacherHandler)

	//Handling the /v1/students. The handler is a type implementing the Handler interface here
	sHandler := studentHandler{}
	mux.Handle("/v1/students", sHandler)

	//Create the server.
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()

}

func teacherHandler(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of teacher's called")
	res.WriteHeader(200)
	res.Write(data)
}

type studentHandler struct{}

func (h studentHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("V1 of student's called")
	res.WriteHeader(200)
	res.Write(data)
}
```

Run the server now

```sh
go run main.go
```

## Must read resources:

- [Beginners guide to serving files using HTTP servers in Go](https://medium.com/rungo/beginners-guide-to-serving-files-using-http-servers-in-go-4e542e628eac)
- [HTTP Server](https://golangr.com/golang-http-server/)
- [ZetCode: Go HTTP server](https://zetcode.com/golang/http-server/)
- [GoDocs: HTTP server in GoLang](https://golangdocs.com/http-server-in-golang)
- [Golang HTTP server for pro](https://scullwm.medium.com/golang-http-server-for-pro-69034c276355)
- [Go Web Examples: HTTP Server](https://gowebexamples.com/http-server/)
