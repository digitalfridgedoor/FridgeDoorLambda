package main

import (
	"context"
	"errors"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

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
	log.Printf("Processing Lambda request ViewRecipe %s\n", request.RequestContext.RequestID)

	recipeID, err := fridgedoorgateway.ReadPathParameterAsObjectID(&request, "id")
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(400), errMissingParameter
	}

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errFind
	}

	r, err := recipeapi.FindOne(context.Background(), recipeID, user)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), errFind
	}

	return fridgedoorgateway.ResponseSuccessful(r), nil
}

func main() {
	lambda.Start(Handler)
}
