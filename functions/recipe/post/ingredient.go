package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

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

	r, err := editable.AddIngredient(context.Background(), request.MethodStepIndex, &ingID)

	return r, err
}

func updateIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.UpdateIngredient(context.Background(), request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeIngredient(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.RemoveIngredient(context.Background(), request.MethodStepIndex, request.IngredientID)

	return r, err
}
