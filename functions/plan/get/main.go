package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/planapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errAuth = errors.New("Auth")
var errBadRequest = errors.New("Bad request")
var errGetResult = errors.New("Error getting result")
var errParse = errors.New("Error parsing response")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request GetPlan %s\n", request.RequestContext.RequestID)

	ok, month, year := parseParameters(&request)
	if !ok {
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}

	isplanninggroup, planningGroupID := tryParsePlanningGroup(&request)

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	ctx := context.TODO()

	if isplanninggroup {
		plan, err := planapi.FindOneForGroup(ctx, *planningGroupID, month, year)
		return planAsResponse(plan, err)
	} else {
		plan, err := planapi.FindOne(ctx, user, month, year)
		return planAsResponse(plan, err)
	}
}

func planAsResponse(plan *dfdmodels.Plan, err error) (events.APIGatewayProxyResponse, error) {
	if err != nil {
		fmt.Printf("Error retrieving plan: %v\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errGetResult
	}

	return fridgedoorgateway.ResponseSuccessful(plan), nil
}

func tryParsePlanningGroup(request *events.APIGatewayProxyRequest) (bool, *primitive.ObjectID) {
	paramValue, ok := request.QueryStringParameters["planningGroupId"]
	if !ok {
		return false, nil
	}

	planningGroupID, err := primitive.ObjectIDFromHex(paramValue)
	if err != nil {
		fmt.Printf("Could not parse query parameter 'planningGroupId' to string, val = '%v'.\n", paramValue)
		return false, nil
	}

	return true, &planningGroupID
}

func parseParameters(request *events.APIGatewayProxyRequest) (bool, int, int) {

	month, ok := tryGetIntQueryParameter(request, "month")
	if !ok {
		return false, 0, 0
	}

	year, ok := tryGetIntQueryParameter(request, "year")
	if !ok {
		return false, 0, 0
	}

	return true, month, year
}

func tryGetIntQueryParameter(request *events.APIGatewayProxyRequest, paramName string) (int, bool) {
	paramValue, ok := request.QueryStringParameters[paramName]
	if !ok {
		fmt.Printf("Missing query parameter '%q'.\n", paramName)
		return -1, false
	}
	i, err := strconv.Atoi(paramValue)
	if err != nil {
		fmt.Printf("Could not parse query parameter '%q' to string, val = '%v'.\n", paramName, paramValue)
		return -1, false
	}

	return i, true
}

func main() {
	lambda.Start(Handler)
}
