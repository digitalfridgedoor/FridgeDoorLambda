package main

import (
	"context"
	"errors"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errFind = errors.New("Error finding results")
var errParseResult = errors.New("Result cannot be parsed")
var errAuth = errors.New("Auth")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request Search clippings %s\n", request.RequestContext.RequestID)

	q, _ := request.QueryStringParameters["q"]

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	clippings, err := clippingapi.SearchByName(context.TODO(), q, user.ViewID, 20)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), errFind
	}

	return fridgedoorgateway.ResponseSuccessful(clippings), nil
}

func main() {
	lambda.Start(Handler)
}
