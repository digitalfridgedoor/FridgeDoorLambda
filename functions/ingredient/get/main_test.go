package main

import (
	"encoding/json"
	"strings"
	"testing"
	"unicode"

	"github.com/digitalfridgedoor/fridgedoorapi"
	"github.com/digitalfridgedoor/fridgedoordatabase/ingredient"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	// Arrange
	apirequest := CreateTestAuthorizedRequest("TestUser")

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	ingredients := []*ingredient.Ingredient{}

	err = json.Unmarshal([]byte(response.Body), &ingredients)
	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Greater(t, len(ingredients), 0)
}

func TestHandlerWithQuery(t *testing.T) {

	// Arrange
	apirequest := CreateTestAuthorizedRequest("TestUser")
	apirequest.QueryStringParameters = make(map[string]string)
	apirequest.QueryStringParameters["q"] = "c"

	// Act
	fridgedoorapi.ConnectOrSkip(t)

	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	ingredients := []*ingredient.Ingredient{}

	err = json.Unmarshal([]byte(response.Body), &ingredients)
	assert.Nil(t, err)
	assert.NotNil(t, ingredients)
	assert.Greater(t, len(ingredients), 0)
	for _, ing := range ingredients {
		startswith := []rune("c")[0]
		assert.True(t, oneWordStartsWith(ing.Name, startswith))
	}
}

func CreateTestAuthorizedRequest(username string) *events.APIGatewayProxyRequest {
	claims := make(map[string]interface{})
	claims["cognito:username"] = username
	authorizer := make(map[string]interface{})
	authorizer["claims"] = claims
	context := events.APIGatewayProxyRequestContext{
		Authorizer: authorizer,
	}
	request := &events.APIGatewayProxyRequest{
		RequestContext: context,
	}

	return request
}

func oneWordStartsWith(ing string, startswith rune) bool {
	words := strings.Fields(ing)
	for _, word := range words {
		if unicode.ToLower([]rune(word)[0]) == startswith {
			return true
		}
	}

	return false
}
