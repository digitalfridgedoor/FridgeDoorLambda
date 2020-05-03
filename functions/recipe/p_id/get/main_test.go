package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

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
	pathParameters["id"] = "5dff9eb2f53f35f9fdcefde2"
	apirequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest("TestUser")
	apirequest.PathParameters = pathParameters

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipe := &dfdmodels.Recipe{}

	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, "5dff9eb2f53f35f9fdcefde2", recipe.ID.Hex())
	assert.Equal(t, "Macho Peas", recipe.Name)
	assert.Equal(t, 1, len(recipe.Method))
	assert.Equal(t, 1, len(recipe.Recipes))

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
