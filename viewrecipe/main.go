package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoordatabase"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var connection fridgedoordatabase.Connection

var errMissingParameter = errors.New("Parameter is missing")
var errFind = errors.New("Cannot find expected entity")
var errParseResult = errors.New("Result cannot be parsed")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request  %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	recipeID, ok := request.PathParameters["id"]
	if !ok || recipeID == "" {
		return events.APIGatewayProxyResponse{}, errMissingParameter
	}

	connection := recipe.New(connection)

	chicken, err := connection.FindOne(context.Background(), recipeID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errFind
	}

	b, err := json.Marshal(chicken)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil

}

func main() {
	connected := connect()
	if connected {
		lambda.Start(Handler)

		connection.Disconnect()
	}

	lambda.Start(Handler)
}

func connect() bool {
	connectionString, err := getConnectionString()

	if err != nil {
		fmt.Printf("Error getting connection string, %v.\n", err)
		return false
	}
	fmt.Printf("Got connection string: len=%v\n", len(connectionString))

	fmt.Printf("Connecting...\n")
	connection = fridgedoordatabase.Connect(context.Background(), connectionString)
	fmt.Printf("Connected.\n")

	return true
}

func getConnectionString() (string, error) {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", err
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	keyname := "mongodb"
	withDecryption := true

	fmt.Println("getting parameter")

	paramOutput, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	fmt.Println("success")

	if err != nil {
		return "", err
	}

	return *paramOutput.Parameter.Value, nil
}
