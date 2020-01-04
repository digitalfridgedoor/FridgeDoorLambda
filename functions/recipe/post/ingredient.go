package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addIngredient(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddIngredient(context.Background(), apiRequest, request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}

func updateIngredient(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" || request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := recipeapi.UpdateIngredient(context.Background(), apiRequest, request.RecipeID, request.MethodStepIndex, request.IngredientID, request.Updates)

	return r, err
}

func removeIngredient(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.IngredientID == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.RemoveIngredient(context.Background(), apiRequest, request.RecipeID, request.MethodStepIndex, request.IngredientID)

	return r, err
}
