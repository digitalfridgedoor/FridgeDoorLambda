package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"
)

func addTag(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	if tag, ok := request.Updates["tag"]; ok {
		r, err := recipeapi.AddTag(context.Background(), apiRequest, request.RecipeID, tag)
		return r, err
	}

	return nil, errMissingProperties
}

func removeTag(ctx context.Context, apiRequest *events.APIGatewayProxyRequest, request *UpdateRecipeRequest) (*recipeapi.Recipe, error) {

	if request.Updates == nil {
		return nil, errMissingProperties
	}

	if tag, ok := request.Updates["tag"]; ok {
		r, err := recipeapi.RemoveTag(context.Background(), apiRequest, request.RecipeID, tag)
		return r, err
	}

	return nil, errMissingProperties
}
