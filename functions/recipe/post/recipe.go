package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func updateRecipe(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	if name, ok := request.Updates["name"]; ok {
		r, err := recipeapi.Rename(context.Background(), apiRequest, request.RecipeID, name)
		return r, err
	}

	if hasImage, ok := request.Updates["hasImage"]; ok {
		b := hasImage == "true"
		r, err := recipeapi.SetImageFlag(context.Background(), apiRequest, request.RecipeID, b)
		return r, err
	}

	return nil, errNoUpdates
}
