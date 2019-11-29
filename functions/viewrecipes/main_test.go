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
	claims["cognito:username"] = "Test"
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
	recipes := []*recipe.Description{}

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Greater(t, len(recipes), 0)
	assert.GreaterOrEqual(t, 25, len(recipes))
}
