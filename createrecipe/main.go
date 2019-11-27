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

var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request  %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	name, ok := request.PathParameters["name"]
	if !ok || name == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400}, errBadRequest
	}

	connection, err := fridgedoorapi.Recipe()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, errServer
	}

	userID, err := primitive.ObjectIDFromHex("5d8f7300a7888700270f7752")

	recipe, err := connection.Create(context.Background(), userID, name)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, errServer
	}

	json, err := json.Marshal(recipe)

	resp := fridgedoorapi.ResponseSuccessful(string(json))
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
