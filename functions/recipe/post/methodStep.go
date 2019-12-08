package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"
)

func addMethodStep(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.Action == "" {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.AddMethodStep(context.Background(), request.RecipeID, request.Action)

	return r, err
}

func updateMethodStep(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := fridgedoorapi.UpdateMethodStep(context.Background(), request.RecipeID, request.MethodStepIndex, request.Updates)

	return r, err
}

func removeMethodStep(ctx context.Context, request *UpdateRecipeRequest) (*recipe.Recipe, error) {

	r, err := fridgedoorapi.RemoveMethodStep(context.Background(), request.RecipeID, request.MethodStepIndex)

	return r, err
}
