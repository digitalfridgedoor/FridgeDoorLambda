package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/digitalfridgedoor/fridgedoorapi/planapi"

	"github.com/digitalfridgedoor/fridgedoorapi"
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

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errAuth
	}

	month, ok := tryGetIntQueryParameter(&request, "month")
	if !ok {
		return events.APIGatewayProxyResponse{}, errBadRequest
	}

	year, ok := tryGetIntQueryParameter(&request, "year")
	if !ok {
		return events.APIGatewayProxyResponse{}, errBadRequest
	}

	plan, err := planapi.FindOne(context.TODO(), user, month, year)
	if err != nil {
		fmt.Printf("Error retrieving plan: %v\n", err)
		return events.APIGatewayProxyResponse{}, errGetResult
	}

	b, err := json.Marshal(plan)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParse
	}

	resp := fridgedoorgateway.ResponseSuccessful(string(b))
	return resp, nil
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
	connected := fridgedoorapi.Connect()
	if connected {
		lambda.Start(Handler)

		fridgedoorapi.Disconnect()
	}
}
