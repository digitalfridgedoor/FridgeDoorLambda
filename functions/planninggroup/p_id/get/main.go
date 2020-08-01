package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/digitalfridgedoor/fridgedoorapi/userviewapi"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/planninggroupapi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errMissingParameter = errors.New("Parameter is missing")
var errFind = errors.New("Cannot find expected entity")
var errParseResult = errors.New("Result cannot be parsed")

// PlanningGroupInfo is the type returned by this handler
type PlanningGroupInfo struct {
	GroupID *primitive.ObjectID `json:"groupID"`
	Name    string              `json:"name"`
	Users   []string            `json:"users"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request ViewPlanningGroup %s\n", request.RequestContext.RequestID)

	groupID, err := fridgedoorgateway.ReadPathParameterAsObjectID(&request, "id")
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(400), errMissingParameter
	}

	ctx := context.TODO()

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(ctx, &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errFind
	}

	r, err := planninggroupapi.FindOne(context.Background(), user, *groupID)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(500), errFind
	}

	users := []string{}

	for _, uid := range r.UserIDs {
		view, err := userviewapi.GetByID(ctx, uid)

		if err != nil {
			fmt.Printf("Error retrieving user, '%v'\n", err)
		} else {
			users = append(users, view.Nickname)
		}
	}

	info := &PlanningGroupInfo{
		GroupID: r.ID,
		Name:    r.Name,
		Users:   users,
	}

	return fridgedoorgateway.ResponseSuccessful(info), nil
}

func main() {
	lambda.Start(Handler)
}
