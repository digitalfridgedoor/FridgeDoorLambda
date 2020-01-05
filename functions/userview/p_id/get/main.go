package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

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
	log.Printf("Processing Lambda request ViewUserView %s\n", request.RequestContext.RequestID)

	viewID, ok := request.PathParameters["id"]
	if !ok || viewID == "" {
		return events.APIGatewayProxyResponse{}, errMissingParameter
	}

	userview, err := userviewapi.GetUserViewByID(context.Background(), viewID)
	if err != nil {
		fmt.Printf("Error getting userview: %v.\n", err)
		return events.APIGatewayProxyResponse{}, errFind
	}

	b, err := json.Marshal(userview)
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