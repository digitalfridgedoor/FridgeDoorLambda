package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errMissingParams = errors.New("Missing parameters")
var errParseResult = errors.New("Result cannot be parsed")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ViewIngredientTree  %s\n", request.RequestContext.RequestID)

	parentID, ok := request.QueryStringParameters["parentID"]
	id, err := primitive.ObjectIDFromHex(parentID)
	if !ok || err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

	ings, err := fridgedoorapi.GetIngredientTree(context.TODO(), id)

	b, err := json.Marshal(ings)
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
