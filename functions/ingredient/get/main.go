package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errParseResult = errors.New("Result cannot be parsed")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request SearchIngredient %s\n", request.RequestContext.RequestID)

	q, _ := request.QueryStringParameters["q"]

	ingredient, err := fridgedoorapi.IngredientCollection(context.TODO())
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), err
	}
	ings, err := ingredient.FindByName(context.TODO(), q)

	if err != nil {
		fmt.Printf("Error searching for ingredients, %v", err)
	}

	fmt.Printf("Searching for %v, got %v results.", q, len(ings))

	return fridgedoorgateway.ResponseSuccessful(ings), nil
}

func main() {
	lambda.Start(Handler)
}
