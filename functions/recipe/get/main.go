package main

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/digitalfridgedoor/fridgedoorapi/search"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

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

	if tagsParam, ok := request.QueryStringParameters["tags"]; ok {
		tags = strings.Split(tagsParam, ",")
	}

	if notTagsParam, ok := request.QueryStringParameters["notTags"]; ok {
		notTags = strings.Split(notTagsParam, ",")
	}

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		log.Println("could not find user")
		return fridgedoorgateway.ResponseUnsuccessful(401), errParseResult
	}

	results, err := search.FindRecipeByTags(context.TODO(), user.ViewID, tags, notTags, 20)
	if err != nil {
		log.Printf("response unsuccessful: %v\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(500), errParseResult
	}

	return fridgedoorgateway.ResponseSuccessful(results), nil
}

func main() {
	lambda.Start(Handler)
}
