package main

import (
	"context"
	"digitalfridgedoor/fridgedoorapi"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

	connection, err := fridgedoorapi.Recipe()

	chicken, err := connection.FindOne(context.Background(), recipeID)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errFind
	}

	b, err := json.Marshal(chicken)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Body = string(b)
	resp.StatusCode = 200
	return resp, nil
}

func main() {
	connected := fridgedoorapi.Connect()
	if connected {
		lambda.Start(Handler)

		fridgedoorapi.Disconnect()
	}

	lambda.Start(Handler)
}
