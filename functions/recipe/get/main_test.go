package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/stretchr/testify/assert"
)

func TestSubstring(t *testing.T) {
	sort := "-name"

	assert.Equal(t, "-", sort[0:1])
}

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetRecipeFindByNameOrTagsPredicate()

	ctx := context.TODO()

	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipeCollection := []*recipeapi.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipeCollection)
	assert.Nil(t, err)
	assert.NotNil(t, recipeCollection)
	assert.Equal(t, len(recipeCollection), 0)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}
