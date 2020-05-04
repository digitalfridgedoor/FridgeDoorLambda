package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.SubRecipeID == nil {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.AddSubRecipe(context.Background(), request.SubRecipeID)

	return r, err
}

func removeSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.SubRecipeID == nil {
		return nil, errMissingProperties
	}

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	r, err := editable.RemoveSubRecipe(context.Background(), request.SubRecipeID)

	return r, err
}
