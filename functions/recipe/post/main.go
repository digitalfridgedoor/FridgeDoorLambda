package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")
var errAuth = errors.New("Auth")

// UpdateRecipeRequest is the expected type for updating recipe
type UpdateRecipeRequest struct {
	RecipeID        *primitive.ObjectID `json:"recipeID"`
	MethodStepIndex int                 `json:"methodStepIndex"`
	IngredientID    string              `json:"ingredientID"`
	SubRecipeID     *primitive.ObjectID `json:"subRecipeID"`
	UpdateType      string              `json:"updateType"`
	Updates         map[string]string   `json:"updates"`
}

var errCannotParse = errors.New("Could not parse request")
var errMissingProperties = errors.New("Request is missing properties")
var errNoUpdates = errors.New("No updates")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	res, err := handleRequest(request)

	if res == nil {
		log.Println("Error caught processing request", err)
		return fridgedoorgateway.ResponseUnsuccessfulString(500, "Unexpected Error"), nil
	}

	return *res, err
}

func handleRequest(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error caught processing request", err)
		}
	}()

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request UpdateRecipe %s\n", request.RequestContext.RequestID)

	r, err := parseRequest(&request)
	if err != nil {
		fmt.Printf("Could not parse body: %v.\n", request.Body)
		return createErrorResponse(400, "Could not parse request")
	}

	if r.RecipeID == nil || r.UpdateType == "" {
		return createErrorResponse(400, "Request is missing properties")
	}

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return createErrorResponse(401, "Unauthorized")
	}

	if r.UpdateType == "R_UPDATE" {
		r, err := updateRecipe(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_ADD" {
		r, err := addMethodStep(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_UPDATE" {
		r, err := updateMethodStep(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_DELETE" {
		r, err := removeMethodStep(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "ING_ADD" {
		r, err := addIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "ING_UPDATE" {
		r, err := updateIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "ING_DELETE" {
		r, err := removeIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_ING_ADD" {
		r, err := addStepIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_ING_UPDATE" {
		r, err := updateStepIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "STEP_ING_DELETE" {
		r, err := removeStepIngredient(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "SUB_ADD" {
		r, err := addSubRecipe(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "SUB_DELETE" {
		r, err := removeSubRecipe(context.Background(), user, r)
		return createResponse(r, err)
	} else if r.UpdateType == "PANIC" {
		panic("oh no!")
	}

	log.Printf("Update type not known %s\n", r.UpdateType)
	return createErrorResponse(400, "Update type not known: '" + r.UpdateType + "'")
}

func findRecipe(ctx context.Context, recipeID *primitive.ObjectID, user *fridgedoorgateway.AuthenticatedUser) (*recipeapi.EditableRecipe, error) {
	return recipeapi.FindOneEditable(context.TODO(), recipeID, user)
}

func parseRequest(request *events.APIGatewayProxyRequest) (*UpdateRecipeRequest, error) {
	r := &UpdateRecipeRequest{}
	err := json.Unmarshal([]byte(request.Body), r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func createResponse(r *recipeapi.Recipe, err error) (*events.APIGatewayProxyResponse, error) {

	if err != nil {
		r := fridgedoorgateway.ResponseUnsuccessful(500)
		return &r, err
	}

	rs := fridgedoorgateway.ResponseSuccessful(r)
	return &rs, nil
}

func createErrorResponse(statusCode int, message string) (*events.APIGatewayProxyResponse, error) {

	response := fridgedoorgateway.ResponseUnsuccessfulString(statusCode, message)

	return &response, errors.New(message)
}

func main() {
	lambda.Start(Handler)
}
