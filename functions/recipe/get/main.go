package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errConnect = errors.New("Cannot connect")
var errFind = errors.New("Cannot find expected entity")
var errParseResult = errors.New("Result cannot be parsed")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ViewRecipes %s\n", request.RequestContext.RequestID)

	tags := []string{}
	notTags := []string{}

	if tagsParam, ok := request.PathParameters["tags"]; ok {
		tags = strings.Split(tagsParam, ",")
	}

	if notTagsParam, ok := request.PathParameters["notTags"]; ok {
		notTags = strings.Split(notTagsParam, ",")
	}

	results, err := recipeapi.FindByTags(context.TODO(), &request, tags, notTags)

	b, err := json.Marshal(results)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

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
