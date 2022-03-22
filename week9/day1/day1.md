## Kafka with golang

![kafka-go](https://repository-images.githubusercontent.com/91830403/cdd95700-fa3e-11e9-91c9-3d18e49f1862)

`kafka-go` provides both low and high level APIs for interacting with Kafka, mirroring concepts and implementing interfaces of the Go standard library to make it easy to use and integrate with existing software.

Comparison with other kafka golang packages:

- <b>sarama</b>, which is by far the most popular but is quite difficult to work with. It is poorly documented, the API exposes low level concepts of the Kafka protocol, and it doesn't support recent Go features like contexts. It also passes all values as pointers which causes large numbers of dynamic memory allocations, more frequent garbage collections, and higher memory usage.

- <b>confluent-kafka-go</b> is a cgo based wrapper around librdkafka, which means it introduces a dependency to a C library on all Go code that uses the package. It has much better documentation than sarama but still lacks support for Go contexts.

- <b>goka</b> is a more recent Kafka client for Go which focuses on a specific usage pattern. It provides abstractions for using Kafka as a message passing bus between services rather than an ordered log of events, but this is not the typical use case of Kafka for us at Segment. The package also depends on sarama for all interactions with Kafka.

#### Creating the Kafka Producer

```go
// the topic and broker address are initialized as constants
const (
	topic          = "message-log"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
```

#### Creating the Kafka Consumer

When creating a consumer, we need to specify it’s group ID. This is because a single topic can have multiple consumers, and each consumers group ID ensures that multiple consumers belonging to the same group ID don’t get repeated messages.

Let’s create another consume function that consumes messages from the Kafka cluster whenever they’re available:

```go
func consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address, broker3Address},
		Topic:   topic,
		GroupID: "my-group",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
```

Now that we’ve defined the functions to send and receive messages, we can put it all together in the main function:

```go
func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}
```

## Must read resources:

- [Sohamkamani: Working with kafka](https://www.sohamkamani.com/golang/working-with-kafka/)
- [kafka golang examples](https://golangexample.com/tag/kafka/)
- [kafka-go ](https://pkg.go.dev/github.com/segmentio/kafka-go#section-documentation)
