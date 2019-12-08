package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

// CreateRecipeRequest is the expected type for updating recipe
type CreateRecipeRequest struct {
	Name       string `json:"name"`
	Collection string `json:"collection"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request  CreateRecipe %s\n", request.RequestContext.RequestID)

	r := &CreateRecipeRequest{}
	err := json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Error attempting to parse body: %v.\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, errBadRequest
	}
	if r.Name == "" || r.Collection == "" {
		fmt.Printf("Missing fields: %v.\n", r)
		return events.APIGatewayProxyResponse{StatusCode: 400}, errBadRequest
	}

	ctx := context.Background()

	recipe, err := fridgedoorapi.CreateRecipe(ctx, &request, r.Collection, r.Name)
	if err != nil {
		fmt.Printf("Error creating recipe: %v.\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, errServer
	}

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, errServer
	}

	json, err := json.Marshal(recipe)

	resp := fridgedoorapi.ResponseSuccessful(string(json))
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