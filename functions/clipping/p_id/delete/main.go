package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

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
	log.Printf("Processing Lambda request ClippingRecipe %s\n", request.RequestContext.RequestID)

	clippingID, err := fridgedoorgateway.ReadPathParameterAsObjectID(&request, "id")
	if err != nil {
		fmt.Printf("Error reading clippingID, %v\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errMissingParameter
	}

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errFind
	}

	err = clippingapi.Delete(context.Background(), user, clippingID)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), errFind
	}

	return fridgedoorgateway.ResponseSuccessfulString(""), nil
}

func main() {
	lambda.Start(Handler)
}
