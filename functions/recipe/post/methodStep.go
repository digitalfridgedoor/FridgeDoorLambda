package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func addMethodStep(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.AddMethodStep(context.Background(), request.RecipeID, request.IngredientID)

	return r, err
}

func removeMethodStep(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.RemoveMethodStep(context.Background(), request.RecipeID, request.MethodStepIndex)

	return r, err
}
