package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddIngredient(context.Background(), user, request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}

func updateIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := recipeapi.UpdateIngredient(context.Background(), user, request.RecipeID, request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveIngredient(context.Background(), user, request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}
