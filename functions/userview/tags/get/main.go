package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errConnect = errors.New("Cannot connect")
var errFind = errors.New("Cannot find expected entity")
var errParseResult = errors.New("Result cannot be parsed")
var errAuth = errors.New("Auth")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ViewUserViews %s\n", request.RequestContext.RequestID)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	view, err := userviewapi.GetUserViewByID(context.Background(), user)
	if err != nil {
		fmt.Printf("Error getting userview: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errConnect
	}

	return fridgedoorgateway.ResponseSuccessful(view.Tags), nil
}

func main() {
	connected := fridgedoorapi.Connect()
	if connected {
		lambda.Start(Handler)

		fridgedoorapi.Disconnect()
	}
}
