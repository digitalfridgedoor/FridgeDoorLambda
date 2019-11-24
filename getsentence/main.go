package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/digitalfridgedoor/fridgedoordatabase"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var connection fridgedoordatabase.Connection

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := getCategories()

	// Let's create the response we'll eventually send, being sure to have CORS headers in place
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Body = string(c)
	resp.StatusCode = 200

	return resp, nil
}

func main() {

	connect()

	lambda.Start(Handler)

	connection.Disconnect()

}

func getConnectionString() string {
	region := "eu-west-2"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(region)},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion(region))
	keyname := "mongodb"
	withDecryption := true

	fmt.Println("getting parameter")

	paramOutput, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	fmt.Println("success")

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return *paramOutput.Parameter.Value
}

func connect() {
	connectionString := getConnectionString() // getEnvironmentVariable("connectionstring")
	fmt.Printf("Got connection string: len=%v\n", len(connectionString))

	fmt.Printf("Connecting...\n")
	connection = fridgedoordatabase.Connect(context.Background(), connectionString)
	fmt.Printf("Connected.\n")
}

func getCategories() []byte {
	fmt.Printf("get categories.\n")
	duration5s, _ := time.ParseDuration("5s")
	findCtx, cancel := context.WithTimeout(context.Background(), duration5s)
	defer cancel()

	parentID, _ := primitive.ObjectIDFromHex("5dac430246ba29343620c1df")
	ingredientConnection := ingredient.New(connection)
	cats := ingredientConnection.IngredientByParentID(findCtx, parentID)
	b, _ := json.Marshal(cats)
	fmt.Printf("end.\n")
	return b
}
