package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	claims := make(map[string]interface{})
	claims["cognito:username"] = "7a401f20-86ca-442f-acf3-20d396c06d33"
	authorizor := make(map[string]interface{})
	authorizor["claims"] = claims

	context := events.APIGatewayProxyRequestContext{
		Authorizer: authorizor,
	}
	apirequest := events.APIGatewayProxyRequest{
		RequestContext: context,
	}
	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "fi"

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(apirequest)

	// Assert
	assert.Nil(t, err)
	recipes := []*recipe.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))
}
