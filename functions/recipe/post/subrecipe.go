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

	r, err := recipeapi.AddSubRecipe(context.Background(), user, request.RecipeID, request.SubRecipeID)

	return r, err
}

func removeSubRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.SubRecipeID == nil {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveSubRecipe(context.Background(), user, request.RecipeID, request.SubRecipeID)

	return r, err
}
