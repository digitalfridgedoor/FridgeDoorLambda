package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errAuth = errors.New("Auth")
var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

type createClippingRequest struct {
	Name string `json:"name"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request CreateRecipe %s\n", request.RequestContext.RequestID)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	r := &createClippingRequest{}
	err = json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Error attempting to parse body: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}

	// todo: validate request

	ctx := context.TODO()

	clippingID, err := clippingapi.Create(ctx, user, r.Name)
	if err != nil {
		fmt.Printf("Error creating clipping: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errBadRequest
	}

	return fridgedoorgateway.ResponseSuccessful(clippingID), nil
}

func main() {
	lambda.Start(Handler)
}
