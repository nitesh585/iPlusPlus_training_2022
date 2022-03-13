## Continue with DynamoDB

### Creating an Amazon DynamoDB Table

```go
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
    "log"
)

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
svc := dynamodb.New(sess)


// Create table Movies
tableName := "Movies"

input := &dynamodb.CreateTableInput{
    AttributeDefinitions: []*dynamodb.AttributeDefinition{
        {
            AttributeName: aws.String("Year"),
            AttributeType: aws.String("N"),
        },
        {
            AttributeName: aws.String("Title"),
            AttributeType: aws.String("S"),
        },
    },
    KeySchema: []*dynamodb.KeySchemaElement{
        {
            AttributeName: aws.String("Year"),
            KeyType:       aws.String("HASH"),
        },
        {
            AttributeName: aws.String("Title"),
            KeyType:       aws.String("RANGE"),
        },
    },
    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(10),
        WriteCapacityUnits: aws.Int64(10),
    },
    TableName: aws.String(tableName),
}

_, err := svc.CreateTable(input)
if err != nil {
    log.Fatalf("Got error calling CreateTable: %s", err)
}

fmt.Println("Created the table", tableName)

```

## Must read resources

- [Amazon DynamoDB Examples Using the AWS SDK for Go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
- [Creating an Amazon DynamoDB Table](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-create-table.html)
- [DynamoDB Go package](https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/)
- [AWS DynamoDB SDK v2 & Golang - Complete Cheat Sheet](https://dynobase.dev/dynamodb-golang-query-examples/)
