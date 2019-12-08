package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func addIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.AddIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}

func updateIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.UpdateIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeIngredient(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.RemoveIngredient(context.Background(), request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}
