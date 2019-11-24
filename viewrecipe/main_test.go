package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/digitalfridgedoor/fridgedoordatabase/recipe"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			expect:  "",
			err:     errMissingParameter,
		},
		{
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     errMissingParameter,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}

func TestHandler(t *testing.T) {

	// Arrange

	pathParameters := make(map[string]string)
	pathParameters["id"] = "5dbc80036eb36874255e7fcd"
	apirequest := events.APIGatewayProxyRequest{PathParameters: pathParameters}

	// Act
	connected := connect()
	if !connected {
		fmt.Println("Cannot connect, skipping test")
		return
	}

	response, err := Handler(apirequest)

	// Assert
	assert.Nil(t, err)
	recipe := &recipe.Recipe{}

	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, "5dbc80036eb36874255e7fcd", recipe.ID.Hex())
	assert.Equal(t, "Nandos chicken", recipe.Name)
	assert.Equal(t, 2, len(recipe.Ingredients))
	assert.Equal(t, 1, len(recipe.Recipes))
}
