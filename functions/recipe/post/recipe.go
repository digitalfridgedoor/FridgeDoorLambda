package main

import (
	"context"
	"errors"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgateway"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func updateRecipe(ctx context.Context, user *fridgedoorgateway.AuthenticatedUser, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	rID, err := primitive.ObjectIDFromHex(request.RecipeID)
	if err != nil {
		return nil, errors.New("Invalid id")
	}

	if name, ok := request.Updates["name"]; ok {
		r, err := recipeapi.Rename(context.Background(), user, &rID, name)
		return r, err
	}

	return recipeapi.UpdateMetadata(context.Background(), user, &rID, request.Updates)
}
