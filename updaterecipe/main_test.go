package main

import (
	"encoding/json"
	"testing"

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
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     errCannotParse,
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
	recipeID := "recipeID"
	updateRequest := &UpdateRecipeRequest{
		RecipeID: recipeID,
	}
	body, err := json.Marshal(updateRequest)
	assert.Nil(t, err)
	apirequest := events.APIGatewayProxyRequest{Body: string(body)}

	// Act
	response, err := Handler(apirequest)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, response.Body, "Got recipe "+recipeID)
}
