package main

import (
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	apirequest := events.APIGatewayProxyRequest{}

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(apirequest)

	// Assert
	assert.Nil(t, err)
	recipes := []*recipe.Description{}

	err = json.Unmarshal([]byte(response.Body), &recipes)
	assert.Nil(t, err)
	assert.NotNil(t, recipes)
	assert.Greater(t, len(recipes), 0)
	assert.GreaterOrEqual(t, 25, len(recipes))
}
