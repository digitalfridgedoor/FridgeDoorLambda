package main

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
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
	ctx := context.Background()
	recipeName := "test-recipe"
	testUserName := "TestUser"
	testUser := fridgedoorgatewaytesting.CreateTestAuthenticatedUser(testUserName)
	recipe, err := recipeapi.CreateRecipe(ctx, testUser, recipeName)
	assert.Nil(t, err)

	r, err := recipeapi.FindOne(ctx, testUser, recipe.ID)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	pathParameters := make(map[string]string)
	pathParameters["id"] = recipe.ID.Hex()
	deleterequest := fridgedoorgatewaytesting.CreateTestAuthorizedRequest(testUserName)
	deleterequest.PathParameters = pathParameters

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*deleterequest)

	// Assert
	assert.Equal(t, 200, response.StatusCode)
	assert.Nil(t, err)

	r, err = recipeapi.FindOne(ctx, testUser, recipe.ID)
	assert.NotNil(t, err)
	assert.Nil(t, r)

	fridgedoorgatewaytesting.DeleteTestUser(testUser)
}
