package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/imageapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errMissingParameters = errors.New("Missing parameters")
var errParseResult = errors.New("Error parsing result")

// ImageURLResponse is the type returned by get image url
type ImageURLResponse struct {
	URL string `json:"url"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	res, err := handleRequest(request)

	if err != nil {
		log.Println("Error caught processing request", err)
	}

	if res == nil {
		log.Println("Response empty, returning unexpected error")
		return fridgedoorgateway.ResponseUnsuccessfulString(500, "Unexpected Error"), nil
	}

	return *res, err
}

func handleRequest(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ImageGet %s\n", request.RequestContext.RequestID)

	verb, ok := request.QueryStringParameters["verb"]
	if !ok {
		fmt.Println("Missing parameter 'verb'.")
		response := fridgedoorgateway.ResponseUnsuccessful(400)
		return &response, errMissingParameters
	}
	key, ok := request.QueryStringParameters["key"]
	if !ok {
		fmt.Println("Missing parameter 'key'.")
		response :=  fridgedoorgateway.ResponseUnsuccessful(400)
		return &response, errMissingParameters
	}

	url, err := imageapi.CreatePresignedURL(verb, key)
	if err != nil {
		log.Fatal(err.Error())
	}

	response := &ImageURLResponse{
		URL: url,
	}

	success := fridgedoorgateway.ResponseSuccessful(response)
	return &success, nil
}

func main() {
	lambda.Start(Handler)
}
