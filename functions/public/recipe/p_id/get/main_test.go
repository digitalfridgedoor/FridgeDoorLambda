package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {

	apirequest := events.APIGatewayProxyRequest{}

	_, err := Handler(apirequest)
	assert.IsType(t, errMissingParameter, err)
}

func TestHandlerCanViewPublicUserRecipe(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	anotherUser := dfdtestingapi.CreateTestAuthenticatedUser("Linked")

	recipeName := "recipe"
	r, err := recipeapi.CreateRecipe(ctx, anotherUser, recipeName)

	editable, err := recipeapi.FindOneEditable(ctx, r.ID, anotherUser)
	assert.Nil(t, err)

	updates := make(map[string]string)
	updates["viewableby_everyone"] = "true"
	editable.UpdateMetadata(ctx, updates)
	assert.Nil(t, err)

	apirequest := &events.APIGatewayProxyRequest{}

	pathParameters := make(map[string]string)
	pathParameters["id"] = r.ID.Hex()
	apirequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipe := &recipeapi.Recipe{}

	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, *r.ID, *recipe.ID)
	assert.Equal(t, recipeName, recipe.Name)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}
