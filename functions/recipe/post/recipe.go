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

	return recipeapi.UpdateMetadata(context.Background(), apiRequest, request.RecipeID, request.Updates)
}
