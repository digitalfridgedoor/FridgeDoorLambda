package main

import (
	"context"
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"github.com/stretchr/testify/assert"
)

func TestHandlerUpdateName(t *testing.T) {

	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindByUsernamePredicate()

	ctx := context.TODO()

	user := dfdtestingapi.CreateTestAuthenticatedUser("TestUser")
	recipeName := "recipe"
	updatedRecipeName := "recipe_updated"

	r, err := recipeapi.CreateRecipe(ctx, user, recipeName)
	assert.Nil(t, err)

	updates := make(map[string]string)
	updates["name"] = updatedRecipeName

	request := &UpdateRecipeRequest{
		RecipeID:   r.ID,
		UpdateType: "R_UPDATE",
		Updates:    updates,
	}

	// Arrange
	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")

	b, err := json.Marshal(request)
	assert.Nil(t, err)
	apirequest.Body = string(b)

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotNil(t, response)

	recipe := &dfdmodels.Recipe{}
	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)

	assert.Equal(t, *r.ID, *recipe.ID)
	assert.Equal(t, updatedRecipeName, recipe.Name)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}

func TestRequestUnmarshalsCorrectly(t *testing.T) {

	o := primitive.NewObjectID()
	body := "{\"recipeId\": \"" + o.Hex() + "\"}"

	// Arrange
	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")
	apirequest.Body = body

	// Act
	r, err := parseRequest(apirequest)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, o, *r.RecipeID)
}
