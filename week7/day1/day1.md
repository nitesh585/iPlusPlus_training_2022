## DynamoDB

DynamoDB allows users to create databases capable of storing and retrieving any amount of data, and serving any amount of traffic. It automatically distributes data and traffic over servers to dynamically manage each customer's requests, and also maintains fast performance.

### Advantages

The two main advantages of DynamoDB are scalability and flexibility. It does not force the use of a particular data source and structure, allowing users to work with virtually anything, but in a uniform way.

Its design also supports a wide range of use from lighter tasks and operations to demanding enterprise functionality. It also allows simple use of multiple languages: Ruby, Java, Python, C#, Erlang, PHP, and Perl.

### Limitations

DynamoDB does suffer from certain limitations, however, these limitations do not necessarily create huge problems or hinder solid development.

You can review them from the following points −

- Capacity Unit Sizes − A read capacity unit is a single consistent read per second for items no larger than 4KB. A write capacity unit is a single write per second for items no bigger than 1KB.

- Provisioned Throughput Min/Max − All tables and global secondary indices have a minimum of one read and one write capacity unit. Maximums depend on region. In the US, 40K read and write remains the cap per table (80K per account), and other regions have a cap of 10K per table with a 20K account cap.

- Provisioned Throughput Increase and Decrease − You can increase this as often as needed, but decreases remain limited to no more than four times daily per table.

- Table Size and Quantity Per Account − Table sizes have no limits, but accounts have a 256 table limit unless you request a higher cap.

- Secondary Indexes Per Table − Five local and five global are permitted.

- Projected Secondary Index Attributes Per Table − DynamoDB allows 20 attributes.

- Partition Key Length and Values − Their minimum length sits at 1 byte, and maximum at 2048 bytes, however, DynamoDB places no limit on values.

- Sort Key Length and Values − Its minimum length stands at 1 byte, and maximum at 1024 bytes, with no limit for values unless its table uses a local secondary index.

- Table and Secondary Index Names − Names must conform to a minimum of 3 characters in length, and a maximum of 255. They use the following characters: AZ, a-z, 0-9, “\_”, “-”, and “.”.

- Attribute Names − One character remains the minimum, and 64KB the maximum, with exceptions for keys and certain attributes.

- Reserved Words − DynamoDB does not prevent the use of reserved words as names.

- Expression Length − Expression strings have a 4KB limit. Attribute expressions have a 255-byte limit. Substitution variables of an expression have a 2MB limit.

### Run DynamoDB Local

#### Using Docker

Make sure you have Docker installed. If you don't have Docker yet, you can get it here.

Open terminal and type:

```sh
docker run -p 8000:8000 amazon/dynamodb-local
And that's pretty much it.
```

Your DynamoDB local instance is now running on port `8000`. If you want to connect to this container using SDK or CLI, don't forget to change the endpoint parameter in the configuration. Otherwise, you'll keep trying to connect to the AWS network.

## Must read resources:

- [How to Install AWS CLI on Ubuntu](https://linoxide.com/how-to-install-aws-cli-on-ubuntu-20-04/#:~:text=1%20Prerequisites%202%20Installation.%20AWS%20CLI%20utility%20package,Contents%20to%20S3%20Bucket.%20...%209%20Conclusion.%20)
- [How to run DynamoDB Local and Offline](https://dynobase.dev/run-dynamodb-locally/)
- [DynamoDB Commands](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/index.html)
- [Step by step guide to configure and query DynamoDB with GoLang](https://medium.com/@spiritualcoder/step-by-step-guide-to-use-dynamodb-with-golang-cd374f159a64)
