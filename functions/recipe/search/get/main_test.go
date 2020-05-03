package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()
	dfdtesting.SetRecipeFindByNamePredicate()

	user, apirequest := fridgedoorgatewaytesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "fi"

	recipeapi.CreateRecipe(context.TODO(), user, "fi_recipe")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipes := []*dfdmodels.Recipe{}

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Equal(t, 1, len(recipes))
}
