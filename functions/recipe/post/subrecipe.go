package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addSubRecipe(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.SubRecipeID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddSubRecipe(context.Background(), apiRequest, request.RecipeID, request.SubRecipeID)

	return r, err
}

func removeSubRecipe(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.SubRecipeID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveSubRecipe(context.Background(), apiRequest, request.RecipeID, request.SubRecipeID)

	return r, err
}
