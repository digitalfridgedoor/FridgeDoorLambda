package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

// CreateIngredientRequest is the expected type for creating new ingredient
type CreateIngredientRequest struct {
	Name string `json:"name"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request CreateRecipe %s\n", request.RequestContext.RequestID)

	r := &CreateIngredientRequest{}
	err := json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Error attempting to parse body: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}
	if r.Name == "" {
		fmt.Printf("Missing fields: %v.\n", r)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}

	ingredient, err := fridgedoorapi.IngredientCollection(context.TODO())
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), err
	}

	ing, err := ingredient.Create(context.TODO(), r.Name)
	if err != nil {
		fmt.Printf("Error creating recipe: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errServer
	}

	return fridgedoorgateway.ResponseSuccessful(ing), nil
}

func main() {
	lambda.Start(Handler)
}
