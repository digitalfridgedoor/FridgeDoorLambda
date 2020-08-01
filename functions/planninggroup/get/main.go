package main

import (
	"context"
	"errors"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/planninggroupapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errAuth = errors.New("Auth")
var errServer = errors.New("Server Error")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request CreateRecipe %s\n", request.RequestContext.RequestID)

	ctx := context.TODO()

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(ctx, &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	groups, err := planninggroupapi.FindAll(ctx, user)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), err
	}

	return fridgedoorgateway.ResponseSuccessful(groups), nil
}

func main() {
	lambda.Start(Handler)
}
