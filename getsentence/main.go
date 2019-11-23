package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	c := connect()

	return events.APIGatewayProxyResponse{
		Body:       string(c),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func getConnectionString() string {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	keyname := "/alias/aws/ssm/mongodb"
	withDecryption := true
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	return param.String()
}

func connect() []byte {
	connectionString := getConnectionString() // getEnvironmentVariable("connectionstring")
	fmt.Printf("Got connection string! %v\n", connectionString)

	databaseCtx := context.Background()

	duration5s, _ := time.ParseDuration("5s")
	findCtx, cancel := context.WithTimeout(databaseCtx, duration5s)
	defer cancel()

	connection := Connect(databaseCtx, connectionString)
	defer connection.Disconnect()

	parentID, _ := primitive.ObjectIDFromHex("5dac430246ba29343620c1df")
	cats := connection.IngredientByParentID(findCtx, parentID)
	b, _ := json.Marshal(cats)
	return b
}

func getEnvironmentVariable(key string) string {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == key {
			return pair[1]
		}
	}

	os.Exit(1)
	return ""
}
