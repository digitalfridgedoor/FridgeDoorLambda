package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

// UpdateRecipeRequest is the expected type for updating recipe
type UpdateRecipeRequest struct {
	RecipeID     string `json:"recipeID"`
	IngredientID string `json:"ingredientID"`
	UpdateType   string `json:"updateType"`
}

var errCannotParse = errors.New("Could not parse request")
var errMissingProperties = errors.New("Request is missing properties")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request %s\n", request.RequestContext.RequestID)

	r := &UpdateRecipeRequest{}
	err := json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Could not parse body: %v.\n", request.Body)
		return events.APIGatewayProxyResponse{StatusCode: 500}, errCannotParse
	}

	if r.RecipeID == "" || r.IngredientID == "" || r.UpdateType == "" {
		return events.APIGatewayProxyResponse{StatusCode: 500}, errMissingProperties
	}

	if r.UpdateType == "ADD" {
		r, err := fridgedoorapi.AddIngredient(context.Background(), r.RecipeID, r.IngredientID)
		return createResponse(r, err)
	} else if r.UpdateType == "DELETE" {
		r, err := fridgedoorapi.RemoveIngredient(context.Background(), r.RecipeID, r.IngredientID)
		return createResponse(r, err)
	}

	return events.APIGatewayProxyResponse{StatusCode: 400}, errors.New("Unknown update type")
}

func createResponse(r *recipe.Recipe, err error) (events.APIGatewayProxyResponse, error) {

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	b, err := json.Marshal(r)

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
