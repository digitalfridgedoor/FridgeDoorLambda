package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addMethodStep(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Action == "" {
		return nil, errMissingProperties
	}

	r, err := recipeapi.AddMethodStep(context.Background(), apiRequest, request.RecipeID, request.Action)

	return r, err
}

func updateMethodStep(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	r, err := recipeapi.UpdateMethodStep(context.Background(), apiRequest, request.RecipeID, request.MethodStepIndex, request.Updates)

	return r, err
}

func removeMethodStep(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	r, err := recipeapi.RemoveMethodStep(context.Background(), apiRequest, request.RecipeID, request.MethodStepIndex)

	return r, err
}
