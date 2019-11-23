package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       "string(c)",
		StatusCode: 200,
	}, nil
}

func main() {

	connect()

	lambda.Start(Handler)
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
	// ptype := "SecureString"
	// value := "testing"
	// output, err := ssmsvc.PutParameter(&ssm.PutParameterInput{
	// 	Name:  &keyname,
	// 	Type:  &ptype,
	// 	Value: &value,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("getting parameter")

	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return param.String()
}

func connect() { //[]byte {
	connectionString := getConnectionString() // getEnvironmentVariable("connectionstring")
	fmt.Printf("Got connection string! len=%v\n", len(connectionString))
	fmt.Printf("Got connection string! %v\n", connectionString[0:8])

	// databaseCtx := context.Background()

	// duration5s, _ := time.ParseDuration("5s")
	// findCtx, cancel := context.WithTimeout(databaseCtx, duration5s)
	// defer cancel()

	// connection := Connect(databaseCtx, connectionString)
	// defer connection.Disconnect()

	// parentID, _ := primitive.ObjectIDFromHex("5dac430246ba29343620c1df")
	// cats := connection.IngredientByParentID(findCtx, parentID)
	// b, _ := json.Marshal(cats)
	// return b
}

func getEnvironmentVariable(key string) string {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == key {
			return pair[1]
		}
	}

	os.Exit(1)
	return ""
}
