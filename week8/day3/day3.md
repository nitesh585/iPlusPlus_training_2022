## Kafka

![Kafka](https://www.andplus.com/hs-fs/hubfs/kafkalogo.jpg?width=2400&height=1342&name=kafkalogo.jpg)

Apache Kafka is a software platform which is based on a distributed streaming process. It is a publish-subscribe messaging system which let exchanging of data between applications, servers, and processors as well. Apache Kafka was originally developed by LinkedIn, and later it was donated to the Apache Software Foundation. Currently, it is maintained by Confluent under Apache Software Foundation. Apache Kafka has resolved the lethargic trouble of data communication between a sender and a receiver.

### How can Kafka help?

#### Publish + Subscribe

At its heart lies the humble, immutable commit log, and from there you can subscribe to it, and publish data to any number of systems or real-time applications. Unlike messaging queues, Kafka is a highly scalable, fault tolerant distributed system, allowing it to be deployed for applications like managing passenger and driver matching at Uber, providing real-time analytics and predictive maintenance for British Gas' smart home, and performing numerous real-time services across all of LinkedIn. This unique performance makes it perfect to scale from one app to company-wide use.

#### Store

An abstraction of a distributed commit log commonly found in distributed databases, Apache Kafka provides durable storage. Kafka can act as a 'source of truth', being able to distribute data across multiple nodes for a highly available deployment within a single data center or across multiple availability zones.

#### Process

An event streaming platform would not be complete without the ability to manipulate that data as it arrives. The Streams API within Apache Kafka is a powerful, lightweight library that allows for on-the-fly processing, letting you aggregate, create windowing parameters, perform joins of data within a stream, and more. Perhaps best of all, it is built as a Java application on top of Kafka, keeping your workflow intact with no extra clusters to maintain.

### installation of kafka

- Step 1 — Creating a User for Kafka

- Step 2 — Downloading and Extracting the Kafka Binaries

```sh
> curl "https://downloads.apache.org/kafka/2.6.3/kafka_2.13-2.6.3.tgz" -o ~/Downloads/kafka.tgz
> mkdir ~/kafka && cd ~/kafka
> tar -xvzf ~/Downloads/kafka.tgz --strip 1
```

- Step 3 — Configuring the Kafka Server

```sh
> nano ~/kafka/config/server.properties
```

Edit the properties file

```sh
# ~/kafka/config/server.properties

delete.topic.enable = true
log.dirs=/home/kafka/logs
```

- Step 4 — Creating Systemd Unit Files and Starting the Kafka Server

zookeeper.service

```sh
> sudo nano /etc/systemd/system/zookeeper.service
```

```bash
# /etc/systemd/system/zookeeper.service
[Unit]
Requires=network.target remote-fs.target
After=network.target remote-fs.target

[Service]
Type=simple
User=kafka
ExecStart=/home/kafka/kafka/bin/zookeeper-server-start.sh /home/kafka/kafka/config/zookeeper.properties
ExecStop=/home/kafka/kafka/bin/zookeeper-server-stop.sh
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
```

kafka.service

```sh
> sudo nano /etc/systemd/system/kafka.service
```

```bash
# /etc/systemd/system/kafka.service
[Unit]
Requires=zookeeper.service
After=zookeeper.service

[Service]
Type=simple
User=kafka
ExecStart=/bin/sh -c '/home/kafka/kafka/bin/kafka-server-start.sh /home/kafka/kafka/config/server.properties > /home/kafka/kafka/kafka.log 2>&1'
ExecStop=/home/kafka/kafka/bin/kafka-server-stop.sh
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
```

Start kafka service

```sh
sudo systemctl start kafka
sudo systemctl enable zookeeper
sudo systemctl enable kafka
```

Kafka server listening on port `9092`.

- Step 5 — Testing the Kafka Installation

Publishing messages in Kafka requires:

A producer, who enables the publication of records and data to topics.
A consumer, who reads messages and data from topics.

create a topic named TutorialTopic:

```sh
~/kafka/bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TutorialTopic
```

Now publish the string "Hello, World" to the TutorialTopic topic:

```sh
echo "Hello, World" | ~/kafka/bin/kafka-console-producer.sh --broker-list localhost:9092 --topic TutorialTopic > /dev/null
```

The following command consumes messages from TutorialTopic. Note the use of the --from-beginning flag, which allows the consumption of messages that were published before the consumer was started:

```sh
~/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic TutorialTopic --from-beginning
```

you will see Hello, World appear in your terminal:

```sh
Output
Hello, World
```

- Step 6 — Hardening the Kafka Server

your installation complete, you can remove the kafka user’s admin privileges. Before you do so, log out and log back in as any other non-root sudo user. If you are still running the same shell session that you started this tutorial with, type exit.

Remove the kafka user from the sudo group:

```sh
sudo deluser kafka sudo
```

## Must read resources:

- [Javatpoint: kafka](https://www.javatpoint.com/apache-kafka)
- [Installation guide](https://www.javatpoint.com/installing-kafka-on-linux)
- [Digitalocean: How To Install Apache Kafka on Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-20-04)
