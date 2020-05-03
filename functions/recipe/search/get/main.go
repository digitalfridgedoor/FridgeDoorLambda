package main

import (
	"context"
	"errors"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

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
	log.Printf("Processing Lambda request SearchRecipes %s\n", request.RequestContext.RequestID)

	q, _ := request.QueryStringParameters["q"]

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	recipes, err := recipeapi.FindByName(context.TODO(), user, q)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), errFind
	}

	return fridgedoorgateway.ResponseSuccessful(recipes), nil
}

func main() {
	connected := fridgedoorapi.Connect()
	if connected {
		lambda.Start(Handler)

		fridgedoorapi.Disconnect()
	}
}
