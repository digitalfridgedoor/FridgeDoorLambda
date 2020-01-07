package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			expect:  "",
			err:     errMissingParameter,
		},
		{
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     errMissingParameter,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}

func TestHandler(t *testing.T) {

	// Arrange
	pathParameters := make(map[string]string)
	pathParameters["id"] = "5de3f416aaef34b1c212f7b7"
	apirequest := dfdtesting.CreateTestAuthorizedRequest("TestUser")
	apirequest.PathParameters = pathParameters

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipe := &recipe.Recipe{}

	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, "5de3f416aaef34b1c212f7b7", recipe.ID.Hex())
	assert.Equal(t, "Roast Dinner", recipe.Name)
	assert.Equal(t, 1, len(recipe.Method))
	assert.Equal(t, 1, len(recipe.Recipes))

	dfdtesting.DeleteUserForRequest(apirequest)
}
