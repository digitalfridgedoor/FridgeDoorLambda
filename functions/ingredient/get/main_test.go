package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"unicode"

	"github.com/digitalfridgedoor/fridgedoorapi/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtesting"
	"github.com/digitalfridgedoor/fridgedoorapi/dfdtestingapi"
	"github.com/digitalfridgedoor/fridgedoorapi/ingredients"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(dfdtesting.FindIngredientByNameTestPredicate)

	ctx := context.TODO()

	ingredient, err := ingredients.IngredientCollection(ctx)
	assert.Nil(t, err)
	ingredient.Create(ctx, "beans")
	ingredient.Create(ctx, "toast")

	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	ingredients := []*dfdmodels.Ingredient{}

	err = json.Unmarshal([]byte(response.Body), &ingredients)
	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Greater(t, len(ingredients), 0)

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}

func TestHandlerWithQuery(t *testing.T) {

	// Arrange
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetIngredientFindPredicate(dfdtesting.FindIngredientByNameTestPredicate)

	ctx := context.TODO()

	apirequest := dfdtestingapi.CreateTestAuthorizedRequest("TestUser")
	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "c"

	ingredient, err := ingredients.IngredientCollection(ctx)
	assert.Nil(t, err)
	ingredient.Create(ctx, "carrots")
	ingredient.Create(ctx, "cream")
	ingredient.Create(ctx, "tomatoes")
	ingredient.Create(ctx, "big crackers")

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	ingredients := []*dfdmodels.Ingredient{}

	err = json.Unmarshal([]byte(response.Body), &ingredients)
	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Equal(t, 3, len(ingredients))
	for _, ing := range ingredients {
		startswith := []rune("c")[0]
		assert.True(t, oneWordStartsWith(ing.Name, startswith))
	}

	dfdtestingapi.DeleteUserForRequest(ctx, apirequest)
}

func oneWordStartsWith(ing string, startswith rune) bool {
	words := strings.Fields(ing)
	for _, word := range words {
		if unicode.ToLower([]rune(word)[0]) == startswith {
			return true
		}
	}

	return false
}
