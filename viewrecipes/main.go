package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errFind = errors.New("Cannot find expected entity")
var errParseResult = errors.New("Result cannot be parsed")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request  %s\n", request.RequestContext.RequestID)

	headers := ""
	for k := range request.Headers {
		headers = k + ","
	}
	log.Printf("Headers  %s\n", headers)
	authorizer := ""

	for k := range request.RequestContext.Authorizer {
		authorizer = k + ","
	}
	log.Printf("authorizer  %s\n", authorizer)

	connection, err := fridgedoorapi.Recipe()
	if err != nil {
		return events.APIGatewayProxyResponse{}, errFind
	}

	recipes, err := connection.List(context.Background())
	if err != nil {
		return events.APIGatewayProxyResponse{}, errFind
	}

	b, err := json.Marshal(recipes)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

	resp := fridgedoorapi.ResponseSuccessful(string(b))
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
