package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/digitalfridgedoor/fridgedoorapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	claims := make(map[string]interface{})
	claims["cognito:username"] = "TestUser"
	authorizor := make(map[string]interface{})
	authorizor["claims"] = claims

	context := events.APIGatewayProxyRequestContext{
		Authorizer: authorizor,
	}
	apirequest := events.APIGatewayProxyRequest{
		RequestContext: context,
	}

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(apirequest)

	// Assert
	assert.Nil(t, err)
	recipeCollection := []*recipeapi.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipeCollection)
	assert.Nil(t, err)
	assert.NotNil(t, recipeCollection)
	assert.Equal(t, len(recipeCollection), 0)
}
