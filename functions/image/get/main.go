package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request SearchIngredient %s\n", request.RequestContext.RequestID)

	verb, ok := request.QueryStringParameters["verb"]
	if !ok {
		fmt.Println("Missing parameter 'verb'.")
		return events.APIGatewayProxyResponse{}, errMissingParameters
	}
	key, ok := request.QueryStringParameters["key"]
	if !ok {
		fmt.Println("Missing parameter 'key'.")
		return events.APIGatewayProxyResponse{}, errMissingParameters
	}

	svc, err := connect()
	if err != nil {
		fmt.Printf("Error connecting: %v.\n", err)
		return events.APIGatewayProxyResponse{}, errMissingParameters
	}

	req, err := createRequest(svc, verb, key)
	if err != nil {
		log.Fatal(err.Error())
	}

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	response := &ImageURLResponse{
		URL: url,
	}

	b, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errParseResult
	}

	resp := fridgedoorgateway.ResponseSuccessful(string(b))
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

func connect() (*s3.S3, error) {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return svc, nil
}

func createRequest(svc *s3.S3, verb string, key string) (*request.Request, error) {
	if verb == "put" {
		req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String("digitalfridgedoorphotos"),
			Key:    aws.String(key),
		})

		return req, nil
	} else if verb == "get" {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("digitalfridgedoorphotos"),
			Key:    aws.String(key),
		})

		return req, nil
	}

	return nil, errors.New("Invalid type")
}
