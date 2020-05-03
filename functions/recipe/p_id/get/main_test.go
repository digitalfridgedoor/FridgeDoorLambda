package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/digitalfridgedoor/fridgedoorapi/recipeapi"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/digitalfridgedoor/fridgedoorapi/fridgedoorgatewaytesting"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdmodels"
	"github.com/digitalfridgedoor/fridgedoordatabase/dfdtesting"

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
	dfdtesting.SetTestCollectionOverride()
	dfdtesting.SetUserViewFindPredicate(func(uv *dfdmodels.UserView, m primitive.M) bool {
		return m["username"] == uv.Username
	})

	user, apirequest := fridgedoorgatewaytesting.CreateTestAuthenticatedUserAndRequest("TestUser")

	recipeName := "recipe"
	r, err := recipeapi.CreateRecipe(context.TODO(), user, recipeName)
	assert.Nil(t, err)

	pathParameters := make(map[string]string)
	pathParameters["id"] = r.ID.Hex()
	apirequest.PathParameters = pathParameters

	// Act
	response, err := Handler(*apirequest)

	// Assert
	assert.Nil(t, err)
	recipe := &dfdmodels.Recipe{}

	err = json.Unmarshal([]byte(response.Body), recipe)
	assert.Nil(t, err)
	assert.NotNil(t, recipe)
	assert.Equal(t, *r.ID, *recipe.ID)
	assert.Equal(t, recipeName, recipe.Name)

	fridgedoorgatewaytesting.DeleteUserForRequest(apirequest)
}
