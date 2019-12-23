package main

import (
	"context"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/userview"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/digitalfridgedoor/fridgedoorapi"

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
	collectionName := "public"
	recipeName := "test-recipe"
	testUser := "test-user"
	request := createTestAuthorizedRequest(testUser)
	recipe, err := recipeapi.CreateRecipe(ctx, request, collectionName, recipeName)
	assert.Nil(t, err)

	recipeID := recipe.ID.Hex()
	r, err := recipeapi.FindOne(ctx, request, recipeID)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	pathParameters := make(map[string]string)
	pathParameters["id"] = recipe.ID.Hex()
	deleterequest := createTestAuthorizedRequest(testUser)
	deleterequest.PathParameters = pathParameters

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*deleterequest)

	// Assert
	assert.Equal(t, 200, response.StatusCode)
	assert.Nil(t, err)

	r, err = recipeapi.FindOne(ctx, request, recipeID)
	assert.NotNil(t, err)
	assert.Nil(t, r)

	userview.Delete(ctx, testUser)
}

func createTestAuthorizedRequest(username string) *events.APIGatewayProxyRequest {
	claims := make(map[string]interface{})
	claims["cognito:username"] = username
	authorizer := make(map[string]interface{})
	authorizer["claims"] = claims
	context := events.APIGatewayProxyRequestContext{
		Authorizer: authorizer,
	}
	request := &events.APIGatewayProxyRequest{
		RequestContext: context,
	}

	return request
}
