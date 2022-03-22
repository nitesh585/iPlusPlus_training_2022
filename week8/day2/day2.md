## Pub/Sub Model

Pub/sub is shorthand for publish/subscribe messaging, an asynchronous communication method in which messages are exchanged between applications without knowing the identity of the sender or recipient.
Pub/Sub messaging helps with this in two crucial ways:

- Allowing developers to create decoupled applications easily with a reliable communication method
- Enabling users to create event-driven architectures easily

![pub-sub-model](https://th.bing.com/th/id/R.9d44c942b4fc2356196a3760d87bde5b?rik=ElhJs3Q%2fJXPCdg&riu=http%3a%2f%2fthenewstack.io%2fwp-content%2fuploads%2f2016%2f04%2fPub-Sub.png&ehk=WuuC4o0oANkL3JRmjC7j7p7IzzjrAYEx4y52CDNY4yc%3d&risl=&pid=ImgRaw&r=0)

### Overview

Four core concepts make up the pub/sub model:

- <b>Topic</b> – An intermediary channel that maintains a list of subscribers to relay messages to that are received from publishers
- <b>Message</b> – Serialized messages sent to a topic by a publisher which has no knowledge of the subscribers
- <b>Publisher</b> – The application that publishes a message to a topic
- <b>Subscriber</b> – An application that registers itself with the desired topic in order to receive the appropriate messages

### Advantages of publish/subscribe pattern

- Decoupled/loosely coupled components
- Decoupled/loosely coupled components
- Real-time communication
- Ease of development
- Increased scalability & reliability
- Testability improvements

### Disadvantages of pub/sub pattern

- Unnecessary complexity in smaller systems
- Media streaming

### How to Use Pub/Sub Messaging

- <b>Parallel Asynchronous Processing</b>
  Events published to a message topic can trigger multiple workers to perform necessary but unrelated tasks at the same time.
- <b>Application and System Alerts</b>
  Instantly deliver critical updates and asynchronous event notifications to affected application components and end-users.
- <b>Manage Workflows</b>
  Relay events among applications, move data between data stores, refresh distributed caches, or update records in business systems.
- <b>Balance Workloads</b>
  Batch up tasks for bulk processing, or break up a larger task into lots of smaller ones and divide the work among multiple workers using message queuing.
- <b>Log to Multiple Systems</b>
  Record events to analyse later, capture logs for operations and security, or archive to meet backup or compliance requirements.
- <b>Use Fanout for Replication</b>
  Replicate data between production and development environments, or enable development and testing using live data.
- <b>Coordinate Serverless Applications</b>
  Fanout asynchronous event notifications to distributed functions and message queues, in order to coordinate the components of your serverless application.
- <b>Stream IoT Data</b>
  Pub/Sub messaging is a very powerful way for IoT devices to interact – devices can easily stream data to backend systems or to each other.

### Pub/sub messaging services

There are multitudes of Pub/Sub messaging services, from dedicated message brokers to cloud offerings. Following is a list of some common Pub/Sub services.

- Apache Kafka. Developed by Apache, Kafka has robust Pub/Sub messaging features with message logs.
- Faye. Simple Pub/Sub service designed to power web applications with servers designed for NodeJS and Ruby.
- Redis. This is one of the most popular message brokers with support for both traditional message queues as well as pub/sub pattern implementations.
- Amazon SNS. The Amazon Simple Notification Service is a fully managed service that offers Pub/Sub messages.
- Google Pub/Sub. GCP offering for pub/sub messaging service implementation.
- Azure Service Bus. A robust messaging service (MaaS) solution that offers Pub/Sub pattern.

## Must read resources

- [Event-Driven Architecture and Pub/Sub Pattern Explained](https://www.altexsoft.com/blog/event-driven-architecture-pub-sub/)
- [What is Pub/Sub Messaging?](https://blog.stackpath.com/pub-sub/)
- [What Is Pub/Sub?](https://www.bmc.com/blogs/pub-sub-publish-subscribe/#:~:text=A%20pub%2Fsub%20model%20allows%20messages%20to%20be%20broadcasted,instantly%20push%20the%20message%20to%20all%20the%20subscribers.)
