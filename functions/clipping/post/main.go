package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/clippingapi"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var errAuth = errors.New("Auth")
var errServer = errors.New("Server Error")
var errBadRequest = errors.New("Bad request")
var errMissingProperties = errors.New("Request is missing properties")

type updateClippingRequest struct {
	ClippingID *primitive.ObjectID `json:"clippingID"`
	UpdateType string              `json:"updateType"`
	Updates    map[string]string   `json:"updates"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing a new Lambda request CreateRecipe %s\n", request.RequestContext.RequestID)

	r := &updateClippingRequest{}
	err := json.Unmarshal([]byte(request.Body), r)
	if err != nil {
		fmt.Printf("Error attempting to parse body: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}

	// todo: validate request

	ctx := context.TODO()

	if r.ClippingID == nil || r.UpdateType == "" {
		return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
	}

	user, err := fridgedoorgateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return fridgedoorgateway.ResponseUnsuccessful(401), errAuth
	}

	if r.UpdateType == "LINK_ADD" {
		name, nameok := r.Updates["linkname"]
		url, urlok := r.Updates["linkurl"]
		if !nameok || !urlok {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		clipping, err := clippingapi.AddLink(ctx, user, r.ClippingID, name, url)
		return response(clipping, err)
	} else if r.UpdateType == "LINK_REMOVE" {
		linkidxs, linkok := r.Updates["linkidx"]
		if !linkok {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		linkidx, err := strconv.Atoi(linkidxs)
		if err != nil {
			return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
		}
		clipping, err := clippingapi.RemoveLink(ctx, user, r.ClippingID, linkidx)
		return response(clipping, err)
	} else if r.UpdateType == "LINK_UPDATE" {
		linkidxs, linkok := r.Updates["linkidx"]
		if !linkok {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		linkidx, err := strconv.Atoi(linkidxs)
		if err != nil {
			return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
		}
		property, propertyok := r.Updates["property"]
		if !propertyok {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		value, valueok := r.Updates["value"]
		if !valueok {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		clipping, err := clippingapi.UpdateLink(ctx, user, r.ClippingID, linkidx, property, value)
		return response(clipping, err)
	} else if r.UpdateType == "LINK_REORDER" {
		linkidxs1, linkok1 := r.Updates["linkidx1"]
		if !linkok1 {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		linkidx1, err := strconv.Atoi(linkidxs1)
		if err != nil {
			return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
		}
		linkidxs2, linkok2 := r.Updates["linkidx2"]
		if !linkok2 {
			return fridgedoorgateway.ResponseUnsuccessful(400), errMissingProperties
		}
		linkidx2, err := strconv.Atoi(linkidxs2)
		if err != nil {
			return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
		}
		clipping, err := clippingapi.SwapLinkPosition(ctx, user, r.ClippingID, linkidx1, linkidx2)
		return response(clipping, err)
	} else if r.UpdateType == "UPDATE" {
		clipping, err := clippingapi.Update(ctx, user, r.ClippingID, r.Updates)
		return response(clipping, err)
	}

	return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
}

func response(clipping *dfdmodels.Clipping, err error) (events.APIGatewayProxyResponse, error) {
	if err != nil {
		fmt.Printf("Error: %v.\n", err)
		return fridgedoorgateway.ResponseUnsuccessful(400), errBadRequest
	}
	return fridgedoorgateway.ResponseSuccessful(clipping), nil
}

func main() {
	lambda.Start(Handler)
}
