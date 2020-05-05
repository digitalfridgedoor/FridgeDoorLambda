package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(dfdtesting.FindIngredientByNameTestPredicate)

	ctx := context.TODO()

	apirequest := dfdtesting.CreateTestAuthorizedRequest("TestUser")

	request := &CreateIngredientRequest{
		Name: "beans",
	}
	b, err := json.Marshal(request)
	assert.Nil(t, err)

	apirequest.Body = string(b)

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	ing := &dfdmodels.Ingredient{}

	err = json.Unmarshal([]byte(response.Body), &ing)
	assert.Nil(t, err)
	assert.NotNil(t, ing)
	assert.Equal(t, "beans", ing.Name)
	assert.NotNil(t, ing.ID)

	ingredient, err := fridgedoorapi.IngredientCollection(ctx)
	assert.Nil(t, err)
	i, err := ingredient.FindOne(ctx, ing.ID)
	assert.Nil(t, err)
	assert.Equal(t, *ing.ID, *i.ID)

	dfdtesting.DeleteUserForRequest(ctx, apirequest)
}
