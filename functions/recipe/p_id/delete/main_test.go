package main

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

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
	dfdtesting.SetTestCollectionOverride()

	ctx := context.Background()
	recipeName := "test-recipe"
	testUserName := "TestUser"
	testUser := dfdtesting.CreateTestAuthenticatedUser(testUserName)
	recipe, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)
	assert.Nil(t, err)

	r, err := recipeapi.FindOne(ctx, recipe.ID, testUser)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	pathParameters := make(map[string]string)
	pathParameters["id"] = recipe.ID.Hex()
	deleterequest := dfdtesting.CreateTestAuthorizedRequest(testUserName)
	deleterequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*deleterequest)

	// Assert
	assert.Equal(t, 200, response.StatusCode)
	assert.Nil(t, err)

	dfdtesting.DeleteTestUser(ctx, testUser)
}
