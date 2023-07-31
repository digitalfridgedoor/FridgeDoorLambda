package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addStepIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	ingID, err := primitive.ObjectIDFromHex(request.IngredientID)
	if err != nil {
		return nil, errBadRequest
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.AddStepIngredient(context.Background(), request.MethodStepIndex, &ingID)

	return r, err
}

func updateStepIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.UpdateStepIngredient(context.Background(), request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeStepIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.RemoveStepIngredient(context.Background(), request.MethodStepIndex, request.IngredientID)

	return r, err
}
