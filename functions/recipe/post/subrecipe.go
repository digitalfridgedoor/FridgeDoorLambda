package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func addSubRecipe(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.SubRecipeID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddSubRecipe(context.Background(), request.RecipeID, request.SubRecipeID)

	return r, err
}

func removeSubRecipe(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.SubRecipeID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveSubRecipe(context.Background(), request.RecipeID, request.SubRecipeID)

	return r, err
}
