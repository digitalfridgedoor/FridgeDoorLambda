package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/planapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errAuth = errors.New("Auth")
var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

type updatePlanRequest struct {
	Month      int                `json:"month"`
	Year       int                `json:"year"`
	Day        int                `json:"day"`
	MealIndex  int                `json:"mealIndex"`
	RecipeName string             `json:"recipeName"`
	RecipeID   primitive.ObjectID `json:"recipeID"`
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

	r := &updatePlanRequest{}
	err = json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Error attempting to parse body: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}
	// todo: validate request

	apirequest := &planapi.UpdateDayPlanRequest{
		Year:       r.Year,
		Month:      r.Month,
		Day:        r.Day,
		MealIndex:  r.MealIndex,
		RecipeName: r.RecipeName,
		RecipeID:   r.RecipeID,
	}

	plan, err := planapi.UpdatePlan(context.TODO(), user, apirequest)
	if err != nil {
		fmt.Printf("Error updating plan: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errServer
	}

	return fridgedoorgateway.ResponseSuccessful(plan), nil
}

func main() {
	lambda.Start(Handler)
}
