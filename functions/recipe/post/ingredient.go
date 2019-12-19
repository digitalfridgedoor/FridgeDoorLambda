package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func addIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}

func updateIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := recipeapi.UpdateIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}
