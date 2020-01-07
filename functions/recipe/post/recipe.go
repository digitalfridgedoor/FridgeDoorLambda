package main

import (
	"context"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func updateRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	if name, ok := request.Updates["name"]; ok {
		r, err := recipeapi.Rename(context.Background(), user, request.RecipeID, name)
		return r, err
	}

	return recipeapi.UpdateMetadata(context.Background(), user, request.RecipeID, request.Updates)
}
