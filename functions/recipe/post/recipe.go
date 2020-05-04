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

	editable, err := findRecipe(ctx, request.RecipeID, user)
	if err != nil {
		return nil, err
	}

	if name, ok := request.Updates["name"]; ok {
		r, err := editable.Rename(context.Background(), name)
		return r, err
	}

	return editable.UpdateMetadata(context.Background(), request.Updates)
}
